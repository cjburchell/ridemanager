export type Gender = 'M' | 'F' | '';

export interface IAthlete {
  id: string;
  strava_athlete_id: number;
  name: string;
  sex: Gender;
  profile: string;
  profile_medium: string;
}

export interface IAchievements {
  first_count: number;
  second_count: number;
  third_count: number;
  finished_count: number;
}
