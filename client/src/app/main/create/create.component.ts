import {Component, Input, OnChanges, OnInit} from '@angular/core';
import {
  ActivityService,
  IActivity, ICategory
} from '../../services/activity.service';
import {IUser} from '../../services/user.service';
import {Router} from '@angular/router';
import * as uuid from 'uuid';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnChanges {

  Activity: IActivity;
  newCategory: ICategory;
  @Input() user: IUser;

  constructor(private activityService: ActivityService, private router: Router) {
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
  }

  clearRoute() {
    this.Activity.route = undefined;
  }


}
