import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'activityTypeToImage'
})
export class ActivityTypeToImagePipe implements PipeTransform {

  transform(value: any, args?: any): any {
    return null;
  }

}
