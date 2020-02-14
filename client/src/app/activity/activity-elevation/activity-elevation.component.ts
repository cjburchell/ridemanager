import {Component, Input, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {IActivity} from '../../services/contracts/activity';
import {ChartDataSets, ChartOptions, ChartType} from 'chart.js';
import {Color, Label} from 'ng2-charts';
import * as mapboxgl from 'mapbox-gl';
import {Polyline} from '../../services/polyline';

@Component({
  selector: 'app-activity-elevation',
  templateUrl: './activity-elevation.component.html',
  styleUrls: ['./activity-elevation.component.scss']
})
export class ActivityElevationComponent implements OnInit, OnChanges {
  lineChartData: ChartDataSets[] = [];

  lineChartLabels: Label[] = undefined;

  lineChartOptions: ChartOptions = {
    responsive: true,
    scales: {
      xAxes: [
        {
          type: 'linear',
          scaleLabel: { display: true, labelString: 'Distance (km)'},
        }
      ],
      yAxes: [
        {
          type: 'linear',
          scaleLabel: { display: true, labelString: 'Elevation (m)'},
        }
      ]
    }
  };

  lineChartColors: Color[] = [];
  lineChartLegend = false;
  lineChartPlugins = [];
  lineChartType: ChartType = 'line';

  @Input() activity: IActivity;
  constructor() { }

  ngOnInit() {
  }

  ngOnChanges(changes: SimpleChanges): void {
    this.lineChartData = [];
    if (this.activity.stages) {
      for (const stage of this.activity.stages) {
        if (stage.elevation) {
          const data = {
            data: stage.elevation,
            backgroundColor: 'red',
            showLine: true,
          };

          this.lineChartData.push(data);
        }
      }
    }

    if (this.activity.route.elevation) {
        const data = {
          data: this.activity.route.elevation,
          backgroundColor: 'blue',
          showLine: true,
        };

        this.lineChartData.push(data);
    }
  }
}
