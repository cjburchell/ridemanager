import {SegmentType} from './activity';


export interface ISegmentSummary {
  id: number;
  name: string;
  activity_type: SegmentType;
  distance: number;
  average_grade: number;
  maximum_grade: number;
  elevation_high: number;
  elevation_low: number;
  climb_category: number;
  start_latlng: number[];
  end_latlng: number[];
  city: string;
  state: string;
  country: string;
  private: boolean;
  starred: boolean;
  map: IMap;
}

export interface IMap {
  id: string;
  polyline: string;
  summary_polyline: string;
}


export interface IRouteSummary {
  id: number;
  name: string;
  distance: number;
  type: number;
  segments: ISegmentSummary[];
  map: IMap;
}
