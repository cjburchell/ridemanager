import { Component, OnInit } from '@angular/core';
import {IAthlete} from '../../services/user.service';

@Component({
  selector: 'app-add-friend-dialog',
  templateUrl: './add-friend-dialog.component.html',
  styleUrls: ['./add-friend-dialog.component.scss']
})
export class AddFriendDialogComponent implements OnInit {
  friendSearchText: any;
  friends: IAthlete[];
  selectedFriend: IAthlete;

  constructor() {
  }

  ngOnInit() {
  }

  selectFriend(friend: IAthlete) {
    this.selectedFriend = friend;
  }

  getFriends() {

  }

  addFriend() {

  }
}
