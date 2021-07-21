import { Component, Input, OnInit, OnDestroy, ChangeDetectorRef } from '@angular/core';
import { LaForgeProvisionStatus, LaForgeSubscribeUpdatedStatusSubscription, LaForgeTeam } from '@graphql';
import { EnvironmentService } from '@services/environment/environment.service';
import { teamChildrenCompleted } from '@util';
import { Subscription } from 'rxjs';

import { RebuildService } from '../../services/rebuild/rebuild.service';

@Component({
  selector: 'app-team',
  templateUrl: './team.component.html',
  styleUrls: ['./team.component.scss']
})
export class TeamComponent implements OnInit, OnDestroy {
  private unsubscribe: Subscription[] = [];
  @Input() title: string;
  @Input() team: LaForgeTeam;
  @Input() style: 'compact' | 'collapsed' | 'expanded';
  @Input() selectable: boolean;
  @Input() mode: 'plan' | 'build' | 'manage';
  isSelectedState = false;
  planStatus: LaForgeSubscribeUpdatedStatusSubscription['updatedStatus'];

  constructor(private rebuild: RebuildService, private envService: EnvironmentService, private cdRef: ChangeDetectorRef) {
    if (!this.mode) this.mode = 'manage';
    if (!this.style) this.style = 'compact';
    if (!this.selectable) this.selectable = false;
  }

  ngOnInit(): void {
    const sub1 = this.envService.statusUpdate.asObservable().subscribe(() => {
      this.checkPlanStatus();
      this.cdRef.detectChanges();
    });
    this.unsubscribe.push(sub1);
  }

  ngOnDestroy() {
    this.unsubscribe.forEach((s) => s.unsubscribe());
  }

  checkPlanStatus(): void {
    this.planStatus = this.envService.getStatus(this.team.TeamToPlan.PlanToStatus.id) || this.planStatus;
  }

  // getStatus(): ProvisionStatus {
  //   // let status: ProvisionStatus = ProvisionStatus.ProvStatusUndefined;
  //   let numWithAgentData = 0;
  //   let totalAgents = 0;
  //   for (const network of this.team.TeamToProvisionedNetwork) {
  //     for (const host of network.ProvisionedNetworkToProvisionedHost) {
  //       totalAgents++;
  //       if (host.ProvisionedHostToAgentStatus?.clientId) numWithAgentData++;
  //     }
  //   }
  //   if (numWithAgentData === totalAgents) return ProvisionStatus.COMPLETE;
  //   else if (numWithAgentData === 0) return ProvisionStatus.FAILED;
  //   else return ProvisionStatus.INPROGRESS;
  // }

  allAgentsResponding(): boolean {
    let numWithAgentData = 0;
    let totalAgents = 0;
    for (const pnet of this.team.TeamToProvisionedNetwork) {
      for (const host of pnet.ProvisionedNetworkToProvisionedHost) {
        totalAgents++;
        if (host.ProvisionedHostToAgentStatus?.clientId) numWithAgentData++;
      }
    }
    return numWithAgentData === totalAgents;
  }

  getStatusIcon(): string {
    if (!this.planStatus) return 'minus-circle';

    switch (this.planStatus.state) {
      case LaForgeProvisionStatus.Planning:
        return 'ruler-triangle fas';
      case LaForgeProvisionStatus.Todelete:
        return 'recycle fas';
      case LaForgeProvisionStatus.Deleteinprogress:
        return 'trash-restore fas';
      case LaForgeProvisionStatus.Deleted:
        return 'trash fas';
      default:
        return 'users';
    }
  }

  getStatusColor(): string {
    if (!this.planStatus) return 'dark';
    switch (this.planStatus.state) {
      case LaForgeProvisionStatus.Complete:
        if (this.allAgentsResponding()) {
          return 'success';
        } else {
          return 'warning';
        }
      case LaForgeProvisionStatus.Inprogress:
        return 'info';
      case LaForgeProvisionStatus.Tainted:
        return 'danger';
      case LaForgeProvisionStatus.Failed:
        return 'danger';
      case LaForgeProvisionStatus.Todelete:
        return 'primary';
      case LaForgeProvisionStatus.Deleteinprogress:
        return 'info';
      case LaForgeProvisionStatus.Planning:
        return 'primary';
      default:
        return 'dark';
    }
  }

  // getStatusColor(): string {
  //   switch (this.getStatus()) {
  //     case ProvisionStatus.COMPLETE:
  //       return 'success';
  //     case ProvisionStatus.INPROGRESS:
  //       return 'warning';
  //     case ProvisionStatus.FAILED:
  //       return 'danger';
  //     default:
  //       return 'dark';
  //   }
  // }

  onSelect(): void {
    let success = false;
    if (!this.isSelected()) {
      success = this.rebuild.addTeam(this.team);
    } else {
      success = this.rebuild.removeTeam(this.team);
    }
    if (success) this.isSelectedState = !this.isSelectedState;
  }

  isSelected(): boolean {
    return this.rebuild.hasTeam(this.team);
  }

  shouldCollapse(): boolean {
    return teamChildrenCompleted(this.team, this.envService.getStatus);
  }
}
