import {Component, Input, OnChanges, ViewChild} from '@angular/core';
import {SelectStageComponent} from '../main/create/select-stage/select-stage.component';
import {SelectRouteComponent} from '../main/create/select-route/select-route.component';
import {IStravaService} from '../services/strava.service';
import {AddCategoryComponent} from '../main/create/add-category/add-category.component';
import {IRouteSummary, ISegmentSummary} from '../services/contracts/strava';
import {IActivity, ICategory, IRoute, IStage} from '../services/contracts/activity';
import {IAthlete} from '../services/contracts/user';


@Component({
  selector: 'app-edit-activity',
  templateUrl: './edit-activity.component.html',
  styleUrls: ['./edit-activity.component.scss']
})
export class EditActivityComponent implements OnChanges {

  @Input() Activity: IActivity;
  sortedStages: IStage[];
  @Input() user: IAthlete;
  @ViewChild('selectStage', {static: false}) selectStage: SelectStageComponent;
  @ViewChild('selectRoute', {static: false}) selectRoute: SelectRouteComponent;
  @ViewChild('addCategoryItem', {static: false}) addCategoryItem: AddCategoryComponent;

  constructor(private stravaService: IStravaService) { }

  ngOnChanges() {
    if (this.Activity !== undefined) {
      this.updateSortedStages();
    }
  }

  private updateSortedStages() {
    this.sortedStages = this.Activity.stages.sort((item1, item2) => item1.number - item2.number);
  }

  deleteCategory(category: ICategory) {
    const index = this.Activity.categories.indexOf(category, 0);
    if (index > -1) {
      this.Activity.categories.splice(index, 1);
    }
  }

  addCategory(newCategory: ICategory) {
    this.Activity.categories.push(newCategory);
  }

  clearRoute() {
    this.Activity.route = undefined;
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

  async addStage(segment: ISegmentSummary) {
    const fullSegment = await this.stravaService.getSegment(segment.id);
    const stage: IStage = {
      segment_id: fullSegment.id,
      distance: fullSegment.distance,
      activity_type: fullSegment.activity_type,
      name: fullSegment.name,
      number: this.Activity.stages.length + 1,
      start_latlng: fullSegment.start_latlng,
      end_latlng: fullSegment.end_latlng,
      map: await this.stravaService.getSegmentMap(segment.id),
    };

    this.Activity.stages.push(stage);
    this.updateDistance();
    this.updateSortedStages();
  }

  async setRoute(selectedRoute: IRouteSummary, addStages: boolean) {
    const route: IRoute = {
      distance: selectedRoute.distance,
      map: await this.stravaService.getRouteMap(selectedRoute.id),
      id: selectedRoute.id,
      name: selectedRoute.name
    };

    this.Activity.route = route;
    if (addStages) {
      const fullRoute = await this.stravaService.getRoute(selectedRoute.id);
      for (const segment of fullRoute.segments) {
        this.addStage(segment);
      }
    }
  }

  private updateDistance() {
    this.Activity.total_distance = this.Activity.stages.reduce((total, item) => total + item.distance, 0);
  }

}
