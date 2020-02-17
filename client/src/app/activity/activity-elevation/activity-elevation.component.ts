import {Component, Input, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {IActivity, IPoint} from '../../services/contracts/activity';
import {ChartDataSets, ChartOptions, ChartType} from 'chart.js';
import {Color, Label} from 'ng2-charts';

@Component({
  selector: 'app-activity-elevation',
  templateUrl: './activity-elevation.component.html',
  styleUrls: ['./activity-elevation.component.scss']
})
export class ActivityElevationComponent implements OnInit, OnChanges {
  constructor() {
  }

  lineChartData: ChartDataSets[] = [];

  lineChartLabels: Label[] = undefined;

  lineChartOptions: ChartOptions = {
    tooltips: {enabled: false},
    elements: {point: {radius: 0, hoverRadius: 0}},
    maintainAspectRatio: false,
    responsive: true,
    scales: {
      xAxes: [
        {
          type: 'linear',
          // scaleLabel: { display: true, labelString: 'Distance (km)'},
          ticks: {
            callback: (value) => {
              return (value / 1000).toFixed(2) + 'km';
            }
          }
        }
      ],
      yAxes: [
        {
          type: 'linear',
          // scaleLabel: { display: true, labelString: 'Elevation (m)'},
          ticks: {
            callback: (value) => {
              return value + 'm';
            }
          }
        }
      ]
    }
  };

  lineChartColors: Color[] = [];
  lineChartLegend = false;
  lineChartPlugins = [];
  lineChartType: ChartType = 'line';

  @Input() activity: IActivity;

  private static findOffset(startPoint: number[], routeMap: IPoint[], offset: number): number {
    const result = [];
    for (let i = 0; i < routeMap.length; i++) {
      if (offset > routeMap[i].d) {
        continue;
      }

      const latDiff = routeMap[i].p[0] - startPoint[0];
      const longDiff = routeMap[i].p[1] - startPoint[1];

      const diff = Math.abs(latDiff) + Math.abs(longDiff);

      result.push({value: diff, index: i});
    }

    const index = result.sort((a, b) => a.value - b.value)[0].index;
    return routeMap[index].d;
  }

  ngOnInit() {
  }

  ngOnChanges(changes: SimpleChanges): void {
    this.lineChartData = [];

    if (this.activity.route.map) {
      let offset = 0;
      for (const stage of this.activity.stages.sort((a, b) => a.number - b.number)) {
        if (stage.map) {
          let maxDistance = 0;
          offset = ActivityElevationComponent.findOffset(stage.start_latlng, this.activity.route.map, offset);
          const stageElevation = [];
          stage.map.forEach((point) => {
            const newPoint = {
              x: point.d + offset,
              y: point.e
            };
            maxDistance = Math.max(newPoint.x, maxDistance);
            stageElevation.push(newPoint);
          });

          const data = {
            data: stageElevation,
            backgroundColor: 'red',
            showLine: true,
          };

          this.lineChartData.push(data);
          offset = maxDistance;
        }
      }
    } else if (this.activity.stages) {
      let offset = 0;
      for (const stage of this.activity.stages.sort((a, b) => a.number - b.number)) {
        if (stage.map) {
          let maxDistance = 0;
          const stageElevation = [];
          stage.map.forEach((point) => {
            const newPoint = {
              x: point.d + offset,
              y: point.e
            };
            maxDistance = Math.max(newPoint.x, maxDistance);
            stageElevation.push(newPoint);
          });

          const data = {
            data: stageElevation,
            backgroundColor: 'red',
            showLine: true,
          };

          this.lineChartData.push(data);
          offset = maxDistance;
        }
      }
    }

    if (this.activity.route.map) {
      const elevation = [];
      this.activity.route.map.forEach((point) => {
        const newPoint = {
          x: point.d,
          y: point.e
        };
        elevation.push(newPoint);
      });

      const data = {
        data: elevation,
        backgroundColor: 'blue',
        showLine: true,
      };

      this.lineChartData.push(data);
    }
  }
}
