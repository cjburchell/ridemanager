import {IActivityService} from '../activity.service';
import {ISettingsService} from '../settings.service';
import {IStravaService} from '../strava.service';
import {IUserService} from '../user.service';
import {ITokenService} from '../token.service';
import {IActivity, IParticipant, IResult} from '../contracts/activity';
import {IRouteSummary, ISegmentSummary} from '../contracts/strava';
import {IAchievements, IAthlete} from '../contracts/user';

const me: IAthlete = {
  id: '8e8ff23c-b3da-46ee-bc08-376d37115af7',
  strava_athlete_id: 10018917,
  name: 'Christiaan Burchell',
  profile: 'https://dgalywyr863hv.cloudfront.net/pictures/athletes/10018917/3021422/6/large.jpg',
  profile_medium: 'https://dgalywyr863hv.cloudfront.net/pictures/athletes/10018917/3021422/6/medium.jpg',
  sex: 'M'
};


const activity1: IActivity = {
  activity_id: 'fc9f764b-20ba-49fa-b3a9-0fcd6989ff0f',
  activity_type: 'group_ride',
  owner: me,
  name: 'test',
  description: '',
  start_time: new Date('2020-01-04T23:15:52.68-05:00'),
  end_time: new Date('2020-01-11T23:15:52.68-05:00'),
  total_distance: 2641.8,
  duration: 604800,
  time_left: -2738112.7186438,
  starts_in: -3342912.7186438,
  route: {
    id: 1,
    name: 'MSM Enduro',
    distance: 8769.18,
    map: {
      // tslint:disable-next-line:max-line-length
      polyline: 'wxkwG|abnMfCwOzBuEQm@a@vAm@a@vByEkAeA\\u@lAIWcBvAmAl@Z[yAj@Ir@c@S}@AiBdAm@e@kBrCdAh@GZaBhByAgD?b@s@uA}BRkA~@oA_@]\\cBb@l@f@EJo@rCNdAY^yA_@u@[Xg@eAjAsB`@wHj@u@BgAk@{AYwGlDcATHZj@TFNPVl@p@F`BYj@[LSx@Wb@FVANJZdAv@Z\\`@RG`@a@d@Fd@Ph@z@FR`@XZb@\\~BEXNl@NVd@PDZr@`@Xd@VFFPf@PPb@Hb@PPPFPGh@l@VDpBfAR?b@a@RFJt@CXIXQJSAs@v@Ir@Ar@UzAVj@Ij@Sl@Iz@Cd@Fj@OpA@nATbC?Gq@PqByJaBvGCsEs@PUxBSa@JeAc@h@Ow@c@zCgBw@_@k@Mh@u@EoA}BwAt@uEK?@}@kGMaBJsIEgAEYQc@KOQISGc@N_A|@?AGlDdAVg@zAt@pAIzAcAVa@vBb@|@m@l@aAfFg@W]|B|@nAI~@mAf@k@jADzAOz@q@o@E~Bc@aAMdBKb@yAQq@hDgA|@Qa@mFpERy@Kk@yApA@_Au@hBm@a@O[z@{B_GnA_FvDgEXMz@_@WK^n@~A}ASEt@yDiFsFuD}@vAc@tA`@m@Dz@s@`AyCaCkAG]dAtHnBpKlDlCfAdChAjHjCNGJfA',
    }
  },
  privacy: 'private',
  categories: [
    {
      category_id: '2e95ceed-6afe-469a-8c21-4f36c744bc82',
      name: 'Open'
    }
  ],
  stages: [
    {
      segment_id: 0,
      distance: 1213,
      number: 1,
      activity_type: 'Ride',
      name: 'Wild-Side Enduro Stage 1 (2016)',
      map: {
        // tslint:disable-next-line:max-line-length
        polyline: 'uliwGdn`nMGGCIMuAKi@A_@UYIOSeAGMSYEMI{@SS]XDRMbAENi@bAEPEb@BP@j@EDO?Os@C]XoAH{@EBIPMd@KP?UDYDMBUKBEFa@nAKr@EHMEBw@J[?[KOKf@An@CJKNCKBc@A[MIGDGPCREl@OrAK@O_@QWIIe@EIIUg@GEEDAd@GDk@FMMCUGOo@iAWCYBQDBP?\\ONW?WEc@SaAUa@QWC',
      },
      start_latlng: [
        45.928592,
        -75.865462
      ],
      end_latlng: [
        45.933615,
        -75.863959
      ]
    },
    {
      segment_id: 1,
      distance: 1428.8,
      number: 2,
      activity_type: 'Ride',
      name: 'WildSide Enduro Stage 2 (2016)',
      map: {
        // tslint:disable-next-line:max-line-length
        polyline: '}sjwG~n_nM@NBP?NBJINCRNPTALDDp@KXDT@j@|@`AV`ACTMVy@@MGMAMBKJM`@Er@Ob@LHX?JBDJALIHODa@HMFKLGLI\\Qd@a@vBGJKAGGGAEDGZAz@@PDLLTLNPHFFBL?d@Kb@IHK@WAIBIJY|@c@z@G`@Af@Kt@GHO[GGEBEFGTBl@CJG@IIM[SIEHJn@Af@I`@ULa@Dc@KKDSZKRGTEh@IXKPc@Za@j@GQFa@OIaBbBGNM?c@RmA`AKBZc@HGNc@OMSEM@',
      },
      start_latlng: [
        45.934873,
        -75.860471
      ],
      end_latlng: [
        45.93931,
        -75.869704
      ]
    }
  ],
  participants: [
    {
      athlete: me,
      category_id: '0',
      results: [{
        rank: 1,
        segment_id: 0,
        time: 50000,
        stage_number: 0,
      }, {
        rank: 1,
        segment_id: 1,
        time: 50000,
        stage_number: 1,
      }],
      time: 100000,
      rank: 1,
      out_of: 1,
      stages: 2,
      offset_time: 0
    }
  ],
  state: 'finished',
  max_participants: 10
};

const activities: IActivity[] = [activity1];

const segments: ISegmentSummary[] = [
  {
    private: false,
    id: 0,
    distance: 1213,
    activity_type: 'Ride',
    name: 'Wild-Side Enduro Stage 1 (2016)',
    map: {
      // tslint:disable-next-line:max-line-length
      polyline: 'uliwGdn`nMGGCIMuAKi@A_@UYIOSeAGMSYEMI{@SS]XDRMbAENi@bAEPEb@BP@j@EDO?Os@C]XoAH{@EBIPMd@KP?UDYDMBUKBEFa@nAKr@EHMEBw@J[?[KOKf@An@CJKNCKBc@A[MIGDGPCREl@OrAK@O_@QWIIe@EIIUg@GEEDAd@GDk@FMMCUGOo@iAWCYBQDBP?\\ONW?WEc@SaAUa@QWC',
    },
    start_latlng: [
      45.928592,
      -75.865462
    ],
    end_latlng: [
      45.933615,
      -75.863959
    ]
  },
  {
    private: false,
    id: 1,
    distance: 1428.8,
    activity_type: 'Ride',
    name: 'WildSide Enduro Stage 2 (2016)',
    map: {
      // tslint:disable-next-line:max-line-length
      polyline: '}sjwG~n_nM@NBP?NBJINCRNPTALDDp@KXDT@j@|@`AV`ACTMVy@@MGMAMBKJM`@Er@Ob@LHX?JBDJALIHODa@HMFKLGLI\\Qd@a@vBGJKAGGGAEDGZAz@@PDLLTLNPHFFBL?d@Kb@IHK@WAIBIJY|@c@z@G`@Af@Kt@GHO[GGEBEFGTBl@CJG@IIM[SIEHJn@Af@I`@ULa@Dc@KKDSZKRGTEh@IXKPc@Za@j@GQFa@OIaBbBGNM?c@RmA`AKBZc@HGNc@OMSEM@',
    },
    start_latlng: [
      45.934873,
      -75.860471
    ],
    end_latlng: [
      45.93931,
      -75.869704
    ]
  }
];

const achevements: IAchievements = {
  finished_count: 10,
  first_count: 1,
  second_count: 3,
  third_count: 6
};

const routes: IRouteSummary[] = [
    {
      distance: 16869.324,
      id: 0,
      map: {
        polyline: undefined,
      },
      name: 'Powerline loop',
      type: 1,
      segments: [],
    },
  {
    segments: [],
    distance: 7017.786,
    id: 1,
    map: {
      polyline: undefined
    },
    name: 'E-Ride 2',
    type: 1,
  },
  {
    segments: [],
    distance: 5312.458,
    id: 2,
    map: {
      polyline: undefined
    },
    name: 'Test Run',
    type: 2,
  },
  {
    segments,
    distance: 8769.18,
    id: 3,
    map: {
      // tslint:disable-next-line:max-line-length
      polyline: 'wxkwG|abnMfCwOzBuEQm@a@vAm@a@vByEkAeA\\u@lAIWcBvAmAl@Z[yAj@Ir@c@S}@AiBdAm@e@kBrCdAh@GZaBhByAgD?b@s@uA}BRkA~@oA_@]\\cBb@l@f@EJo@rCNdAY^yA_@u@[Xg@eAjAsB`@wHj@u@BgAk@{AYwGlDcATHZj@TFNPVl@p@F`BYj@[LSx@Wb@FVANJZdAv@Z\\`@RG`@a@d@Fd@Ph@z@FR`@XZb@\\~BEXNl@NVd@PDZr@`@Xd@VFFPf@PPb@Hb@PPPFPGh@l@VDpBfAR?b@a@RFJt@CXIXQJSAs@v@Ir@Ar@UzAVj@Ij@Sl@Iz@Cd@Fj@OpA@nATbC?Gq@PqByJaBvGCsEs@PUxBSa@JeAc@h@Ow@c@zCgBw@_@k@Mh@u@EoA}BwAt@uEK?@}@kGMaBJsIEgAEYQc@KOQISGc@N_A|@?AGlDdAVg@zAt@pAIzAcAVa@vBb@|@m@l@aAfFg@W]|B|@nAI~@mAf@k@jADzAOz@q@o@E~Bc@aAMdBKb@yAQq@hDgA|@Qa@mFpERy@Kk@yApA@_Au@hBm@a@O[z@{B_GnA_FvDgEXMz@_@WK^n@~A}ASEt@yDiFsFuD}@vAc@tA`@m@Dz@s@`AyCaCkAG]dAtHnBpKlDlCfAdChAjHjCNGJfA'
    },
    name: 'MSM Enduro',
    type: 1,
  },
  {
    segments: [],
    distance: 8460.628,
    id: 4,
    map: {
      polyline: undefined
    },
    name: 'E- Ride',
    type: 1,
  }
];

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
    return new Promise<string>(resolve => {
      switch (setting) {
        case 'mapboxAccessToken':
          resolve('pk.eyJ1IjoiY2pidXJjaGVsbCIsImEiOiJjaXJzc2hpNDMwaTY0ZmZtNnViZzg5NmRrIn0.LSuKYxyhwBDpHCtjim-g0A');
          break;
        default:
          resolve('test');
      }

    });
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
