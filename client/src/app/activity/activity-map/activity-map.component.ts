import {Component, Input, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {IActivity} from '../../services/activity.service';
import * as mapboxgl from 'mapbox-gl';
import {LngLatLike, LngLatBoundsLike} from 'mapbox-gl';
import {Polyline} from '../../services/polyline';
import * as geojson from 'geojson';
import {SettingsService} from '../../services/settings.service';

@Component({
  selector: 'app-activity-map',
  templateUrl: './activity-map.component.html',
  styleUrls: ['./activity-map.component.scss']
})
export class ActivityMapComponent implements OnInit, OnChanges {

  map: mapboxgl.Map;
  style = 'mapbox://styles/mapbox/outdoors-v11';
  @Input() activity: IActivity;
  private token: string;

  private static swapLatLong(points: number[][]): number[][] {
    for (const point of points) {
      const temp = point[0];
      point[0] = point[1];
      point[1] = temp;
    }
    return points;
  }

  constructor(private settingsService: SettingsService) {
  }

  async ngOnInit() {
    this.settingsService.getSetting('mapboxAccessToken').subscribe(token => {
      this.token = token;
      // @ts-ignore
      mapboxgl.accessToken = token;

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
    });
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

        const points = Polyline.decode(stage.map.polyline);
        for (const point of points) {
          maxLat = Math.max(maxLat, point[0]);
          minLat = Math.min(minLat, point[0]);
          maxLong = Math.max(maxLong, point[1]);
          minLong = Math.min(minLong, point[1]);
        }
      }
    }

    if (this.activity.route) {
      if (this.activity.route.map && this.activity.route.map.polyline) {
        const decodedPoints = Polyline.decode(this.activity.route.map.polyline);
        for (const point of decodedPoints) {
          maxLat = Math.max(maxLat, point[0]);
          minLat = Math.min(minLat, point[0]);
          maxLong = Math.max(maxLong, point[1]);
          minLong = Math.min(minLong, point[1]);
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
                coordinates: ActivityMapComponent.swapLatLong(Polyline.decode(this.activity.route.map.polyline)),
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

        const tempMap = this.map;
        const icon = '/assets/images/bike.png';
        this.map.loadImage(icon, (error, image) => {
          if (error) {
            throw error;
          }
          tempMap.addImage('icon', image);
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
                      coordinates: ActivityMapComponent.swapLatLong(Polyline.decode(stage.map.polyline)),
                    },
                    properties: {},
                  }
              },
              layout: {

              },
              paint: {
                'line-color': '#F00',
                'line-width': 3,
              }
            };
            tempMap.addLayer(trackLayer);

            const pointslayer: mapboxgl.Layer = {
              id: 'stage_points' + stage.number,
              type: 'line',
              source: {
                type: 'geojson',
                data: {
                  type: 'FeatureCollection',
                  features: [{
                      type: 'Feature',
                      geometry: {
                        type: 'Point',
                        coordinates: [stage.start_latlng[1], stage.start_latlng[0]],
                      },
                      properties: {
                        title: 'Start Stage ' + stage.number,
                        icon: 'monument',
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
                        title: 'Start Stage ' + stage.number,
                        description: stage.name,
                        url: 'https://www.strava.com/segments/' + stage.segment_id
                      },
                    }
                  ]
                }
              },
              layout: {
              }
            };
            tempMap.addLayer(pointslayer);
          }
        });
      }
    });
  }
}
