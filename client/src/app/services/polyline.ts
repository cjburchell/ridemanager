
import * as polyline from '@mapbox/polyline';

export class Polyline {
  public static decode(str: string): number[][] {
    return polyline.decode(str);
  }
}
