import {ChangeDetectorRef, Component, EventEmitter, OnInit, Output} from '@angular/core';
import {ISegmentSummary, StravaService} from '../../../services/strava.service';


@Component({
  selector: 'app-select-stage',
  templateUrl: './select-stage.component.html',
  styleUrls: ['./select-stage.component.scss']
})
export class SelectStageComponent implements OnInit {

  stageSearchText = '';
  loadingStages: boolean;
  stages: ISegmentSummary[];
  selectedStage: ISegmentSummary;
  @Output() stageSelected: EventEmitter<ISegmentSummary> = new EventEmitter();

  constructor(private stravaService: StravaService,
              private ref: ChangeDetectorRef) { }

  ngOnInit() {
   this.show();
  }

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
    this.loadingStages = true;
    this.stravaService.getStaredSegments().subscribe((segments: ISegmentSummary[]) => {
      this.stages = segments.filter(item => !item.private);
      this.loadingStages = false;
      this.ref.detectChanges();
    });
  }

}
