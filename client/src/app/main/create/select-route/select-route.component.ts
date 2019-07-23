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
  loading: boolean;
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
    this.autoAddStages = true;
    this.selectedRoute = undefined;
    this.routes = undefined;
    this.getRoutes();
  }

  getRoutes() {
    this.loading = true;
    const perPage = 100;
    this.routes = [];
    const loop = (page: number) => {
      this.stravaService.getRoutes(page, perPage).subscribe((routes: IRouteSummary[]) => {
        this.routes.concat(routes);
        if (routes.length !== perPage) {
          this.loading = false;
          this.ref.detectChanges();
        } else {
          loop(page + 1);
        }

      });
    };

    loop(0);
  }

  selectRoute(route: IRouteSummary) {
    this.selectedRoute = route;
  }
}
