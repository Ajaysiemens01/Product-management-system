import { Component } from '@angular/core';
import { HomeComponent } from './home/home.component';
import { ExcelService } from './excel.service';
import { Product } from './product';

@Component({
  selector: 'app-root',
  imports: [HomeComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'my-app';
   
}
