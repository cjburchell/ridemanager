import {Component, Input, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {IActivity} from '../../services/activity.service';
import * as mapboxgl from 'mapbox-gl';
import { environment } from '../../../environments/environment';
import {LngLatLike} from 'mapbox-gl';
import {LngLatBoundsLike} from 'mapbox-gl';
import {polyline} from '@mapbox/polyline';

@Component({
  selector: 'app-activity-map',
  templateUrl: './activity-map.component.html',
  styleUrls: ['./activity-map.component.scss']
})
export class ActivityMapComponent implements OnInit, OnChanges {
  map: mapboxgl.Map;
  style = 'mapbox://styles/mapbox/outdoors-v11';
  lat = 37.75;
  lng = -122.41;
  center: LngLatLike = [this.lng, this.lat];
  boundingBox: LngLatBoundsLike;
  @Input() activity: IActivity;

  constructor() {
  }

  ngOnInit() {
    // @ts-ignore
    mapboxgl.accessToken = environment.mapbox.accessToken;
    this.map = new mapboxgl.Map({
      container: 'map',
      style: this.style,
      zoom: 13,
      center: this.center
    });

    // Add map controls
    this.map.addControl(new mapboxgl.NavigationControl());
  }

  ngOnChanges(changes: SimpleChanges): void {
    let maxLat = -180;
    let minLat = 180;
    let maxLong = -180;
    let minLong = 180;

    if(this.activity.route.map && this.activity.route.map.polyline)
    {
      let points = polyline.decode(this.activity.route.map.polyline);

      for (let point in points){
        maxLat = Math.max(maxLat, point[0]);
        minLat = Math.min(minLat, point[0]);
        maxLong = Math.max(maxLong, point[1]);
        minLong = Math.min(minLong, point[1]);
      }
    }




    this.map.setCenter(this.center);
    this.map.fitBounds(this.boundingBox);
    if (this.activity.route.map) {
      const polylineRouteOptions = {
        color: '#00F'
      };

      polyline(this.activity.route.map.polyline, polylineRouteOptions).addTo(this.map);
    }
  }
}
