import {Routes} from '@angular/router';
import {HomeComponent} from './home/home.component';
import { ReportComponent } from './report/report.component';
const routeConfig: Routes = [
    {
      path: '',
      component: HomeComponent,
      title: 'Home page',
    },
    {
      path: 'report',
      component: ReportComponent,
      title: 'Home details',
    },
  ];

export default routeConfig;
