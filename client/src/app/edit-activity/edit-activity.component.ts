import {Component, Input, OnChanges, ViewChild} from '@angular/core';
import { IActivity, ICategory, IStage} from '../services/activity.service';
import {IAthlete} from '../services/user.service';
import {SelectStageComponent} from '../main/create/select-stage/select-stage.component';
import {SelectRouteComponent} from '../main/create/select-route/select-route.component';
import {IRouteSummary, ISegmentSummary, StravaService} from '../services/strava.service';
import {AddCategoryComponent} from '../main/create/add-category/add-category.component';


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

  constructor(private stravaService: StravaService) { }

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
