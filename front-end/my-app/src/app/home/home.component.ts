import { Component, inject } from '@angular/core';
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
export class HomeComponent { 

  products: Product[] = [];
    
  excelService: ExcelService = inject(ExcelService);
  constructor() {
    this.loadProducts();
  }

  async loadProducts() {
    this.excelService.fetchProducts(); // Fetches data immediately on service load
    this.excelService.products$.subscribe((data: Product[]) => {
      this.products=data
    });
  }
  

}

