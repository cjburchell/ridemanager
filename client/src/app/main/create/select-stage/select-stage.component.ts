import {ChangeDetectorRef, Component, EventEmitter, OnInit, Output} from '@angular/core';
import {ISegmentSummary, StravaService} from '../../../services/strava.service';


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

  constructor(private stravaService: StravaService,
              private ref: ChangeDetectorRef) { }

  public show() {
    this.stageSearchText = '';
    this.selectedStage = undefined;
    this.stages = undefined;
    this.getStages();
  }

  selectStage(stage: ISegmentSummary) {
    this.selectedStage = stage;
  }

  getStages() {
    this.loading = true;
    const perPage = 100;
    this.stages = [];
    const loop = (page: number) => {
      this.stravaService.getStaredSegments(page, perPage).subscribe((segments: ISegmentSummary[]) => {
        this.stages = this.stages.concat(segments.filter(item => !item.private));
        if (segments.length !== perPage) {
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
