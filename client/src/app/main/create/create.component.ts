import {ChangeDetectorRef, Component, Input, OnChanges} from '@angular/core';
import {
  ActivityService,
  IActivity, ICategory, IStage
} from '../../services/activity.service';
import {IUser} from '../../services/user.service';
import {Router} from '@angular/router';
import * as uuid from 'uuid';
import {IRouteSummary, ISegmentSummary, StravaService} from '../../services/strava.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnChanges {

  Activity: IActivity;
  newCategory: ICategory;
  @Input() user: IUser;

  stageSearchText = '';
  routeSearchText = '';
  loadingRouts: boolean;
  loadingStages: boolean;
  stages: ISegmentSummary[];
  selectedStage: ISegmentSummary;
  autoAddStages: boolean;

  routes: IRouteSummary[];
  selectedRoute: IRouteSummary;

  constructor(private activityService: ActivityService,
              private router: Router,
              private stravaService: StravaService,
              private ref: ChangeDetectorRef) {
  }

  ngOnChanges() {
    if (this.user !== undefined) {
      this.Activity = {
        activity_id: undefined,
        activity_type: 'group_ride',
        owner_id: this.user.id,
        name: undefined,
        description: undefined,
        start_time: new Date(),
        end_time: new Date(),
        total_distance: undefined,
        duration: undefined,
        time_left: undefined,
        starts_in: undefined,
        route: undefined,
        privacy: 'private',
        categories: [],
        stages: [],
        participants: [],
        state: 'upcoming',
        max_participants: 10
      };

      this.Activity.end_time.setDate(this.Activity.end_time.getDate() + 7);
    }
  }

  back() {
    this.router.navigate([`/main`]);
  }

  create() {
    // TODO: add validation
    this.activityService.createActivity(this.Activity);
  }

  deleteCategory(category: ICategory) {
    const index = this.Activity.categories.indexOf(category, 0);
    if (index > -1) {
      this.Activity.categories.splice(index, 1);
    }
  }

  showAddCategory() {
    this.newCategory = {
      category_id: uuid.v4().toString(),
      name: undefined,
    };
  }

  addCategory(newCategory: ICategory) {
    this.Activity.categories.push(newCategory);
  }

  showRoutes() {
    this.routeSearchText = '';
    this.selectedRoute = undefined;
    this.routes = undefined;
    this.getRoutes();
  }

  clearRoute() {
    this.Activity.route = undefined;
  }


  showAddStage() {
    this.stageSearchText = '';
    this.selectedStage = undefined;
    this.stages = undefined;
    this.getStages();
  }

  deleteStage(stage: IStage) {
    const index = this.Activity.stages.indexOf(stage, 0);
    if (index > -1) {
      this.Activity.stages.splice(index, 1);
    }
  }

  moveStageUp(stage: IStage) {
    // TODO reorder stages
  }

  moveStageDown(stage: IStage) {
    // TODO reorder stages
  }

  selectStage(stage: ISegmentSummary) {
    this.selectedStage = stage;
  }

  getStages() {
    this.loadingStages = true;
    this.stravaService.getStaredSegments().subscribe((segments: ISegmentSummary[]) => {
      this.stages = segments.filter(item => !item.private);
      this.loadingStages = false;
      this.ref.detectChanges();
    });
  }

  addStage(segment: ISegmentSummary) {
    this.stravaService.getSegment(segment.id).subscribe( (fullSegment: ISegmentSummary) => {
        this.Activity.stages.push({
          segment_id: fullSegment.id,
          distance: fullSegment.distance,
          activity_type: fullSegment.activity_type,
          name: fullSegment.name,
          number: this.Activity.stages.length,
          map: fullSegment.map,
          start_latlng: fullSegment.start_latlng,
          end_latlng: fullSegment.end_latlng
        });
    });
  }

  getRoutes() {
    this.loadingRouts = true;
    this.stravaService.getRoutes().subscribe((routes: IRouteSummary[]) => {
      this.routes = routes;
      this.loadingRouts = false;
      this.ref.detectChanges();
    });
  }

  setRoute(selectedRoute: IRouteSummary, addStages: boolean) {
    this.Activity.route = selectedRoute;
    this.stravaService.getRoute(selectedRoute.id).subscribe(
      (fullRoute: IRouteSummary) => {
        this.Activity.route.map = fullRoute.map;
        if (addStages) {
          for (const segment of fullRoute.segments) {
            this.addStage(segment);
          }
        }
      }
    );
  }

  selectRoute(route: IRouteSummary) {
    this.selectedRoute = route;
  }
}
