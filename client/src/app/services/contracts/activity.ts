import {IMap} from './strava';
import {IAthlete} from './user';

export type ActivityType = 'group_ride' | 'race' | 'triathlon' | 'group_run' | 'group_ski';
export type ActivityState = 'upcoming' |'in_progress' |'finished';
export type ActivityPrivacy = 'public' | 'private';

export type SegmentType =
  'Ride'
  | 'AlpineSki'
  | 'BackcountrySki'
  | 'Hike'
  | 'IceSkate'
  | 'InlineSkate'
  | 'NordicSki'
  | 'RollerSki'
  | 'Run'
  | 'Walk'
  | 'Workout'
  | 'Snowboard'
  | 'Snowshoe'
  | 'Kitesurf'
  | 'Windsurf'
  | 'Swim'
  | 'VirtualRide'
  | 'EBikeRide'
  | 'WaterSport'
  | 'Canoeing'
  | 'Kayaking'
  | 'Rowing'
  | 'StandUpPaddling'
  | 'Surfing'
  | 'Crossfit'
  | 'Elliptical'
  | 'RockClimbing'
  | 'StairStepper'
  | 'WeightTraining'
  | 'Yoga'
  | 'WinterSport'
  | 'CrossCountrySkiing';


export interface IActivity {
  activity_id: string;
  activity_type: ActivityType;
  owner: IAthlete;
  name: string;
  description: string;
  start_time: Date;
  end_time: Date;
  total_distance: number;
  duration: number;
  time_left: number;
  starts_in: number;
  route: IRoute;
  privacy: ActivityPrivacy;
  categories: ICategory[];
  stages: IStage[];
  participants: IParticipant[];
  state: ActivityState;
  max_participants: number;
}

export interface IRoute {
  id: number;
  name: string;
  distance: number;
  map: IMap;
}

export interface IParticipant {
  athlete: IAthlete;
  category_id: string;
  results: IResult[];
  time: number;
  rank: number;
  out_of: number;
  stages: number;
  offset_time: number;
}
export interface IResult {
  rank: number;
  segment_id: number;
  activity_id: number;
  time: number;
  stage_number: number;
}

export interface ICategory {
  category_id: string;
  name: string;
}

export interface IStage {
  segment_id: number;
  distance: number;
  number: number;
  activity_type: SegmentType;
  name: string;
  map: IMap;
  start_latlng: number[];
  end_latlng: number[];
}
