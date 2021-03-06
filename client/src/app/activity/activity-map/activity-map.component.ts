import {Component, Input, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import * as mapboxgl from 'mapbox-gl';
import {LngLatBoundsLike, LngLatLike} from 'mapbox-gl';
import * as geojson from 'geojson';
import {ISettingsService} from '../../services/settings.service';
import {IActivity, IPoint} from '../../services/contracts/activity';

@Component({
  selector: 'app-activity-map',
  templateUrl: './activity-map.component.html',
  styleUrls: ['./activity-map.component.scss']
})
export class ActivityMapComponent implements OnInit, OnChanges {

  map: mapboxgl.Map;
  style = 'mapbox://styles/mapbox/outdoors-v11';
  @Input() activity: IActivity;

  private static swapLatLong(map: IPoint[]): number[][] {
    const points: number[][] = [];
    for (const point of map) {
      points.push([point.p[1], point.p[0]]);
    }
    return points;
  }

  constructor(private settingsService: ISettingsService) {
  }

  async ngOnInit(): Promise<void> {

    // @ts-ignore
    mapboxgl.accessToken = await this.settingsService.getSetting('mapboxAccessToken');

    this.map = new mapboxgl.Map({
      container: 'map',
      style: this.style,
      zoom: 13,
      center: [0, 0]
    });

    // Add map controls
    this.map.addControl(new mapboxgl.NavigationControl());

    if (this.activity) {
      this.UpdateActivity();
    }
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (this.map) {
      this.UpdateActivity();
    }
  }

  private UpdateActivity() {
    let maxLat = -180;
    let minLat = 180;
    let maxLong = -180;
    let minLong = 180;

    if (this.activity.stages) {
      for (const stage of this.activity.stages) {
        maxLat = Math.max(maxLat, stage.start_latlng[0]);
        maxLat = Math.max(maxLat, stage.end_latlng[0]);

        minLat = Math.min(minLat, stage.start_latlng[0]);
        minLat = Math.min(minLat, stage.end_latlng[0]);

        maxLong = Math.max(maxLong, stage.start_latlng[1]);
        maxLong = Math.max(maxLong, stage.end_latlng[1]);

        minLong = Math.min(minLong, stage.start_latlng[1]);
        minLong = Math.min(minLong, stage.end_latlng[1]);

        for (const point of stage.map) {
          maxLat = Math.max(maxLat, point.p[0]);
          minLat = Math.min(minLat, point.p[0]);
          maxLong = Math.max(maxLong, point.p[1]);
          minLong = Math.min(minLong, point.p[1]);
        }
      }
    }

    if (this.activity.route) {
      if (this.activity.route.map) {
        for (const point of this.activity.route.map) {
          maxLat = Math.max(maxLat, point.p[0]);
          minLat = Math.min(minLat, point.p[0]);
          maxLong = Math.max(maxLong, point.p[1]);
          minLong = Math.min(minLong, point.p[1]);
        }
      }
    }

    const centerLat = (maxLat - minLat) / 2 + minLat;
    const centerLong = (maxLong - minLong) / 2 + minLong;
    const center: LngLatLike = [centerLong, centerLat];
    const boundingBox: LngLatBoundsLike = [[
      minLong,
      minLat
    ], [
      maxLong,
      maxLat
    ]];

    this.map.setCenter(center);
    this.map.fitBounds(boundingBox);

    this.map.on('load', () => {
      if (this.activity.route.map) {
        const layer: mapboxgl.Layer = {
          id: 'route',
          type: 'line',
          source: {
            type: 'geojson',
            data: {
              type: 'Feature',
              geometry: {
                type: 'LineString',
                coordinates: ActivityMapComponent.swapLatLong(this.activity.route.map),
              },
              properties: {},
            }
          },
          layout: {
            'line-cap': 'round',
            'line-join': 'round'
          },
          paint: {
            'line-color': '#00F',
            'line-width': 3
          }
        };
        this.map.addLayer(layer);
      }

      if (this.activity.stages) {
          for (const stage of this.activity.stages) {
            const trackLayer: mapboxgl.Layer = {
              id: 'stage' + stage.number,
              type: 'line',
              source: {
                type: 'geojson',
                data: {
                  type: 'Feature',
                  geometry: {
                    type: 'LineString',
                    coordinates: ActivityMapComponent.swapLatLong(stage.map),
                  },
                  properties: {},
                }
              },
              layout: {},
              paint: {
                'line-color': '#F00',
                'line-width': 3,
              }
            };
            this.map.addLayer(trackLayer);

            const layout: mapboxgl.SymbolLayout = {
              'icon-image': ['concat', ['get', 'icon'], '-15'],
              'text-field': ['get', 'title'],
              'text-font': ['Open Sans Semibold', 'Arial Unicode MS Bold'],
              'text-offset': [0, 0.6],
              'text-anchor': 'top'
            };

            function stageIcon(): string {
              switch (stage.activity_type) {
                case 'Swim':
                  return 'swimming';
                case 'Ride':
                  return 'bicycle';
                default:
                  return 'triangle';
              }
            }

            const pointsLayer: mapboxgl.Layer = {
              id: 'stage_points' + stage.number,
              type: 'symbol',
              source: {
                type: 'geojson',
                data: {
                  type: 'FeatureCollection',
                  features: [
                    {
                    type: 'Feature',
                    geometry: {
                      type: 'Point',
                      coordinates: [stage.start_latlng[1], stage.start_latlng[0]],
                    },
                    properties: {
                      title: 'Start Stage ' + stage.number,
                      icon: stageIcon(),
                      description: stage.name,
                      url: 'https://www.strava.com/segments/' + stage.segment_id
                    },
                  }, {
                    type: 'Feature',
                    geometry: {
                      type: 'Point',
                      coordinates: [stage.end_latlng[1], stage.end_latlng[0]],
                    },
                    properties: {
                      title: 'End Stage ' + stage.number,
                      icon: 'embassy',
                      description: stage.name,
                      url: 'https://www.strava.com/segments/' + stage.segment_id
                    },
                  }
                  ]
                }
              },
              layout
            };
            this.map.addLayer(pointsLayer);
          }
      }
    });
  }
}
