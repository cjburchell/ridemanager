import {IActivityService} from '../activity.service';
import {ISettingsService} from '../settings.service';
import {IStravaService} from '../strava.service';
import {IUserService} from '../user.service';
import {ITokenService} from '../token.service';
import {IActivity, IParticipant} from '../contracts/activity';
import {IRouteSummary, ISegmentSummary} from '../contracts/strava';
import {IAchievements, IAthlete} from '../contracts/user';

const me: IAthlete = {
  id: 'test',
  name: 'test',
  profile: '',
  profile_medium: '',
  sex: 'M',
  strava_athlete_id: 0
};

const activities: IActivity[] = [{
  activity_id: '0',
  activity_type: 'group_ride',
  categories: [{
    category_id: '0',
    name: 'Open'
  }],
  description: 'This is a test',
  duration: 50000000,
  end_time: new Date(),
  max_participants: 10,
  name: 'test activity',
  owner: me,
  participants: [{
    athlete: me,
    category_id: '0',
    offset_time: 0,
    out_of: 1,
    rank: 1,
    results: [{
      rank: 1,
      segment_id: 0,
      activity_id: 0,
      time: 500000,
      stage_number: 0,
    }],
    stages: 1,
    time: 0
  }],
  privacy: 'private',
  route: undefined,
  stages: [{
    segment_id: 0,
    distance: 5000,
    number: 0,
    activity_type: 'Ride',
    name: 'test',
    map: undefined,
    start_latlng: [],
    end_latlng: [],
  }],
  start_time: new Date(),
  starts_in: 0,
  state: 'upcoming',
  time_left: 0,
  total_distance: 10000
}];


const segments: ISegmentSummary[] = [
  {
    activity_type: 'Ride',
    average_grade: 0,
    city: '',
    climb_category: 0,
    country: '',
    distance: 5000,
    elevation_high: 0,
    elevation_low: 0,
    end_latlng: [],
    id: 0,
    map: undefined,
    maximum_grade: 0,
    name: 'test1',
    private: false,
    starred: false,
    start_latlng: [],
    state: ''
  },
];

const achevements: IAchievements = {
  finished_count: 10,
  first_count: 1,
  second_count: 3,
  third_count: 6
};

const routes: IRouteSummary[] = [{
  distance: 10000, id: 0, map: undefined, name: 'test route', segments, type: 0
}];

export class MockDataService implements IActivityService, ISettingsService, IStravaService, IUserService, ITokenService {
  addParticipant(activity: IActivity, participant: IParticipant): Promise<boolean> {
    activity.participants.push(participant);
    return new Promise<boolean>(resolve => resolve(true));
  }

  checkLogin(): Promise<boolean> {
    return new Promise<boolean>(resolve => resolve(true));
  }

  createActivity(activity: IActivity): Promise<string> {
    activities.push(activity);
    return new Promise<string>(resolve => resolve('test'));
  }

  deleteActivity(activity: IActivity): Promise<any> {
    const index = activities.indexOf(activity, 0);
    if (index > -1) {
      activities.splice(index, 1);
    }

    return new Promise<any>(resolve => resolve('ok'));
  }

  getAchievements(): Promise<IAchievements> {
    return new Promise<IAchievements>(resolve => resolve(achevements));
  }

  getActivities(): Promise<IActivity[]> {
    return new Promise<IActivity[]>(resolve => resolve(activities));
  }

  getActivity(activityId: string): Promise<IActivity> {
    return new Promise<IActivity>(resolve => resolve(activities[0]));
  }

  getJoined(): Promise<IActivity[]> {
    return new Promise<IActivity[]>(resolve => resolve(activities));
  }

  getMe(): Promise<IAthlete> {
    return new Promise<IAthlete>(resolve => resolve(me));
  }

  getMyActivities(): Promise<IActivity[]> {
    return new Promise<IActivity[]>(resolve => resolve(activities));
  }

  getRoute(routeId: number): Promise<IRouteSummary> {
    return new Promise<IRouteSummary>(resolve => resolve(routes[routeId]));
  }

  getRoutes(page: number, perPage: number): Promise<IRouteSummary[]> {
    return new Promise<IRouteSummary[]>(resolve => resolve(routes));
  }

  getSegment(segmentId: number): Promise<ISegmentSummary> {
    return new Promise<ISegmentSummary>(resolve => resolve(segments[segmentId]));
  }

  getSetting(setting: string): Promise<string> {
    return new Promise<string>(resolve => resolve('test'));
  }

  getStaredSegments(page: number, perPage: number): Promise<ISegmentSummary[]> {
    return new Promise<ISegmentSummary[]>(resolve => resolve(segments));
  }

  getToken(): string {
    return 'token';
  }

  leaveActivity(activity: IActivity, athleteId: string): Promise<boolean> {
    return new Promise<boolean>(resolve => resolve(true));
  }

  logOut() {
  }

  setToken(token: string) {
  }

  updateActivity(activity: IActivity): Promise<any> {
    return new Promise<boolean>(resolve => resolve(true));
  }

  updateResults(activity: IActivity): Promise<boolean> {
    return new Promise<boolean>(resolve => resolve(true));
  }

  updateUserResults(activity: IActivity, athleteId: string): Promise<boolean> {
    return new Promise<boolean>(resolve => resolve(true));
  }

  validateToken(): Promise<boolean> {
    return new Promise<boolean>(resolve => resolve(true));
  }
}
