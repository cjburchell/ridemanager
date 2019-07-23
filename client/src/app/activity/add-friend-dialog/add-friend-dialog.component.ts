import {ChangeDetectorRef, Component, EventEmitter, OnInit, Output} from '@angular/core';
import {IAthlete} from '../../services/user.service';
import {IRouteSummary, StravaService} from '../../services/strava.service';

@Component({
  selector: 'app-add-friend-dialog',
  templateUrl: './add-friend-dialog.component.html',
  styleUrls: ['./add-friend-dialog.component.scss']
})
export class AddFriendDialogComponent implements OnInit {
  friendSearchText: any;
  friends: IAthlete[];
  selectedFriend: IAthlete;
  loading: boolean;

  @Output() addFriend: EventEmitter<IAthlete>;

  constructor(private stravaService: StravaService,
              private ref: ChangeDetectorRef) {
  }

  show() {
    this.friendSearchText = '';
    this.selectedFriend = undefined;
    this.friends = undefined;
    this.getFriends();
  }

  ngOnInit() {
    this.show();
  }

  selectFriend(friend: IAthlete) {
    this.selectedFriend = friend;
  }

  getFriends() {
    this.loading = true;
    const perPage = 100;
    this.friends = [];
    const loop = (page: number) => {
      this.stravaService.getFriends(page, perPage).subscribe((friends: IAthlete[]) => {
        this.friends.concat(friends);
        if (friends.length !== perPage) {
          this.loading = false;
          this.ref.detectChanges();
        } else {
          loop(page + 1);
        }

      });
    };

    loop(0);
  }
}
