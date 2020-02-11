import {ChangeDetectorRef, Component, EventEmitter, Output} from '@angular/core';
import {ISegmentSummary} from '../../../services/contracts/strava';
import {IStravaService} from '../../../services/strava.service';


@Component({
  selector: 'app-select-stage',
  templateUrl: './select-stage.component.html',
  styleUrls: ['./select-stage.component.scss']
})
export class SelectStageComponent {

  stageSearchText = '';
  loading: boolean;
  stages: ISegmentSummary[];
  selectedStage: ISegmentSummary;
  @Output() stageSelected: EventEmitter<ISegmentSummary> = new EventEmitter();

  constructor(private stravaService: IStravaService,
              private ref: ChangeDetectorRef) {
  }

  public async show() {
    this.stageSearchText = '';
    this.selectedStage = undefined;
    this.stages = undefined;
    await this.getStages();
  }

  selectStage(stage: ISegmentSummary) {
    this.selectedStage = stage;
  }

  async getStages() {
    this.loading = true;
    const perPage = 100;
    this.stages = [];
    const loop = async (page: number) => {
      const segments = await this.stravaService.getStaredSegments(page, perPage);
      this.stages = this.stages.concat(segments.filter(item => !item.private));
      if (segments.length !== perPage) {
        this.loading = false;
        this.ref.detectChanges();
      } else {
        await loop(page + 1);
      }
    };

    await loop(0);
  }
}
