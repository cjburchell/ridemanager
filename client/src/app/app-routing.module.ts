import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from './login/login.component';
import {MainComponent} from './main/main.component';
import {TokenComponent} from './token/token.component';
import {ActivityComponent} from './activity/activity.component';

const routes: Routes = [
  {path: 'main', component: MainComponent },
  {path: 'main/:pageId', component: MainComponent },
  {path: 'activity/:activityId', component: ActivityComponent },
  {path: 'login', component: LoginComponent },
  {path: 'token', component: TokenComponent },
  {path: '', pathMatch: 'full', redirectTo: 'login'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
