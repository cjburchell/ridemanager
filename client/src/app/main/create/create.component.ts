import { Component, Input, OnChanges, ViewChild} from '@angular/core';
import {
  ActivityService,
  IActivity, ICategory, IStage
} from '../../services/activity.service';
import {IAthlete} from '../../services/user.service';
import {Router} from '@angular/router';

import {IRouteSummary, ISegmentSummary, StravaService} from '../../services/strava.service';
import {SelectStageComponent} from './select-stage/select-stage.component';
import {SelectRouteComponent} from './select-route/select-route.component';
import {TokenService} from "../../services/token.service";

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnChanges {

  Activity: IActivity;
  sortedStages: IStage[];
  @Input() user: IAthlete;
  @ViewChild('selectStage', {static: false}) selectStage: SelectStageComponent;
  @ViewChild('selectRoute', {static: false}) selectRoute: SelectRouteComponent;

  constructor(private activityService: ActivityService,
              private router: Router,
              private stravaService: StravaService,
              private tokenService: TokenService) {
  }

  private updateSortedStages() {
    this.sortedStages = this.Activity.stages.sort((item1, item2) => item1.number - item2.number);
  }

  ngOnChanges() {
    if (this.user !== undefined) {
      this.Activity = {
        activity_id: undefined,
        activity_type: 'group_ride',
        owner: this.user,
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
      this.updateSortedStages();
    }
  }

  back() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main`]);
  }

  create() {
    // TODO: add validation
    this.activityService.createActivity(this.Activity).subscribe(result => {
      if (result !== undefined && result !== null) {
        this.tokenService.checkLogin();
        this.router.navigate([`/main`]);
      }
    }, error1 => console.log(error1));
  }

  deleteCategory(category: ICategory) {
    const index = this.Activity.categories.indexOf(category, 0);
    if (index > -1) {
      this.Activity.categories.splice(index, 1);
    }
  }

  showAddCategory() {
  }

  addCategory(newCategory: ICategory) {
    this.Activity.categories.push(newCategory);
  }

  showRoutes() {
    this.selectRoute.show();
  }

  clearRoute() {
    this.Activity.route = undefined;
  }

  showAddStage() {
    this.selectStage.show();
  }

  deleteStage(stage: IStage) {
    const index = this.Activity.stages.indexOf(stage, 0);
    if (index > -1) {
      this.Activity.stages.splice(index, 1);
      this.updateDistance();

      let stageNumber = 1;
      this.Activity.stages.sort((item1, item2) => item1.number - item2.number).forEach(item => {
        item.number = stageNumber;
        stageNumber++;
      });

      this.updateSortedStages();
    }
  }

  moveStageUp(stage: IStage) {
    const otherStage = this.Activity.stages.find((item) => {
      return item.number === (stage.number - 1);
    });

    if (otherStage) {
      otherStage.number = stage.number;
      stage.number--;
      this.updateSortedStages();
    }
  }

  moveStageDown(stage: IStage) {
    const otherStage = this.Activity.stages.find((item) => {
      return item.number === (stage.number + 1);
    });

    if (otherStage) {
      otherStage.number = stage.number;
      stage.number++;
      this.updateSortedStages();
    }
  }

  addStage(segment: ISegmentSummary) {
    this.stravaService.getSegment(segment.id).subscribe((fullSegment: ISegmentSummary) => {
      this.Activity.stages.push({
        segment_id: fullSegment.id,
        distance: fullSegment.distance,
        activity_type: fullSegment.activity_type,
        name: fullSegment.name,
        number: this.Activity.stages.length + 1,
        map: fullSegment.map,
        start_latlng: fullSegment.start_latlng,
        end_latlng: fullSegment.end_latlng
      });
      this.updateDistance();
      this.updateSortedStages();
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

  private updateDistance() {
    this.Activity.total_distance = this.Activity.stages.reduce((total, item) => total + item.distance, 0);
  }
}
