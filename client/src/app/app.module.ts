
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { MainComponent } from './main/main.component';
import { TokenComponent } from './token/token.component';
import { HttpClientModule } from '@angular/common/http';
import { FooterComponent } from './main/footer/footer.component';
import { HeaderComponent } from './main/header/header.component';
import { ActionsComponent  } from './main/main-menu/actions/actions.component';
import { HelpComponent } from './main/main-menu/help/help.component';
import { ResultSummaryComponent } from './main/main-menu/result-summary/result-summary.component';
import { MainMenuComponent } from './main/main-menu/main-menu.component';
import { FaqComponent } from './main/faq/faq.component';
import { AboutComponent } from './main/about/about.component';
import { ManageComponent } from './main/manage/manage.component';
import { HistoryComponent } from './main/history/history.component';
import { CreateComponent } from './main/create/create.component';
import { JoinComponent } from './main/join/join.component';
import { SummeryComponent } from './main/summery/summery.component';
import {FormsModule} from '@angular/forms';
import { OwlDateTimeModule, OwlNativeDateTimeModule } from 'ng-pick-datetime';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { ActivityTypeToImagePipe } from './pipes/activity-type-to-image.pipe';
import { StageTypeToImagePipe } from './pipes/stage-type-to-image.pipe';
import { RouteTypeToIconPipe } from './pipes/route-type-to-icon.pipe';
import { AddCategoryComponent } from './main/create/add-category/add-category.component';
import { SelectStageComponent } from './main/create/select-stage/select-stage.component';
import { SelectRouteComponent } from './main/create/select-route/select-route.component';
import { ActivityListComponent } from './main/activity-list/activity-list.component';
import { ActivityPanelComponent } from './main/summery/activity-pannel/activity-panel.component';
import { SecondsToTimePipe } from './pipes/seconds-to-time.pipe';
import { SecondsToCountdownPipe } from './pipes/seconds-to-countdown.pipe';
import { RankToPanelTypePipe } from './pipes/rank-to-panel-type.pipe';
import { ActivityComponent } from './activity/activity.component';
import { AthleteComponent } from './common/athlete/athlete.component';
import { JoinDialogComponent } from './activity/join-dialog/join-dialog.component';
import { LoginButtonComponent } from './login/login-button/login-button.component';
import { SearchComponent } from './common/search/search.component';
import { FilterPipe } from './pipes/filter.pipe';
import { LoadingComponent } from './common/loading/loading.component';
import { EditComponent } from './edit/edit.component';
import { EditActivityComponent } from './edit-activity/edit-activity.component';
import { ChartsModule } from 'ng2-charts';
import { DateTimeComponent } from './common/date-time/date-time.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    MainComponent,
    TokenComponent,
    FooterComponent,
    HeaderComponent,
    ActionsComponent,
    HelpComponent,
    ResultSummaryComponent,
    MainMenuComponent,
    FaqComponent,
    AboutComponent,
    ManageComponent,
    HistoryComponent,
    CreateComponent,
    JoinComponent,
    SummeryComponent,
    ActivityTypeToImagePipe,
    StageTypeToImagePipe,
    RouteTypeToIconPipe,
    AddCategoryComponent,
    SelectStageComponent,
    SelectRouteComponent,
    ActivityListComponent,
    ActivityPanelComponent,
    SecondsToTimePipe,
    SecondsToCountdownPipe,
    RankToPanelTypePipe,
    ActivityComponent,
    AthleteComponent,
    JoinDialogComponent,
    LoginButtonComponent,
    SearchComponent,
    FilterPipe,
    LoadingComponent,
    EditComponent,
    EditActivityComponent,
    DateTimeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    OwlDateTimeModule,
    OwlNativeDateTimeModule,
    BrowserAnimationsModule,
    ChartsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
