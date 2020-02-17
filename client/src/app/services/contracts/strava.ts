import {SegmentType} from './activity';


export interface ISegmentSummary {
  private: boolean;
  id: number;
  name: string;
  activity_type: SegmentType;
  distance: number;
  start_latlng: number[];
  end_latlng: number[];
}

export interface IRouteSummary {
  id: number;
  name: string;
  distance: number;
  type: number;
  segments: ISegmentSummary[];
}
