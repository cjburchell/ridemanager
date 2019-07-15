import {ChangeDetectorRef, Component, EventEmitter, OnInit, Output} from '@angular/core';
import {IRouteSummary, ISegmentSummary, StravaService} from '../../../services/strava.service';

export interface IRouteSet {
  addStages: boolean;
  route: IRouteSummary;
}

@Component({
  selector: 'app-select-route',
  templateUrl: './select-route.component.html',
  styleUrls: ['./select-route.component.scss']
})
export class SelectRouteComponent implements OnInit {

  routeSearchText = '';
  loadingRouts: boolean;
  autoAddStages: boolean;
  routes: IRouteSummary[];
  selectedRoute: IRouteSummary;

  @Output() routeSelected: EventEmitter<IRouteSet> = new EventEmitter();

  constructor(private stravaService: StravaService,
              private ref: ChangeDetectorRef) {
  }

  ngOnInit() {
    this.show();
  }

  show() {
    this.routeSearchText = '';
    this.selectedRoute = undefined;
    this.routes = undefined;
    this.getRoutes();
  }

  getRoutes() {
    this.loadingRouts = true;
    this.stravaService.getRoutes().subscribe((routes: IRouteSummary[]) => {
      this.routes = routes;
      this.loadingRouts = false;
      this.ref.detectChanges();
    });
  }

  selectRoute(route: IRouteSummary) {
    this.selectedRoute = route;
  }
}
