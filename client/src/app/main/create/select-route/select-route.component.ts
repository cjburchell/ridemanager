import {ChangeDetectorRef, Component, EventEmitter, Output} from '@angular/core';
import {IStravaService} from '../../../services/strava.service';
import {IRouteSummary} from '../../../services/contracts/strava';

export interface IRouteSet {
  addStages: boolean;
  route: IRouteSummary;
}

@Component({
  selector: 'app-select-route',
  templateUrl: './select-route.component.html',
  styleUrls: ['./select-route.component.scss']
})
export class SelectRouteComponent {

  routeSearchText = '';
  loading: boolean;
  autoAddStages: boolean;
  routes: IRouteSummary[];
  selectedRoute: IRouteSummary;

  @Output() routeSelected: EventEmitter<IRouteSet> = new EventEmitter();

  constructor(private stravaService: IStravaService,
              private ref: ChangeDetectorRef) {
  }

  async show() {
    this.routeSearchText = '';
    this.autoAddStages = true;
    this.selectedRoute = undefined;
    this.routes = undefined;
    await this.getRoutes();
  }

  async getRoutes() {
    this.loading = true;
    const perPage = 100;
    this.routes = [];
    const loop = async (page: number) => {
      const routes = await this.stravaService.getRoutes(page, perPage);
      this.routes = this.routes.concat(routes);
      if (routes.length !== perPage) {
        this.loading = false;
        this.ref.detectChanges();
      } else {
        await loop(page + 1);
      }
    };

    await loop(0);
  }

  selectRoute(route: IRouteSummary) {
    this.selectedRoute = route;
  }
}
