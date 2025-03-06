import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductComponent } from '../product/product.component';
import { Product } from '../product';
import { ExcelService } from '../excel.service';

@Component({
  selector: 'app-home',
  imports: [CommonModule, ProductComponent],
  template: `
  <section class="content">
      <button class="primary" type="button">Add Product</button>
      <button class="primary" type="button">Report</button>
  </section>
  <section class="results">
    <app-product *ngFor="let product of products" [product]="product"></app-product>
  </section>
  `,
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit{

  products: Product[] = [];
    
  excelService: ExcelService = inject(ExcelService);

  constructor() {}

  ngOnInit() {
    console.log("AppComponent Initialized");
    this.excelService.fetchProducts(); // Fetching the products only once
  
    this.excelService.products$.subscribe((data: Product[]) => {
      console.log(" Received Products:", data);
      this.products = data;  // Now properly updating
    });
  }
}
