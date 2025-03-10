import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductComponent } from '../product/product.component';
import { Product } from '../product';
import { ExcelService } from '../excel.service';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, ProductComponent, FormsModule], 
  template: `
  <section class="content">
    <button class="primary" type="button" (click)="toggleModal()">Add Product</button>
    <button class="primary" type="button" (click)="navigateToReport()">Report</button>
  </section>

  <!-- Modal for Adding Product -->
  <div class="modal-overlay" *ngIf="isModalOpen" (click)="toggleModal()">
      <div class="modal" (click)="$event.stopPropagation()">
          <div class="modal-content">
              <h2>Add Product</h2>

              <label>Name:</label>
              <input type="text" [(ngModel)]="newProduct.Name">

              <label>Description:</label>
              <input type="text" [(ngModel)]="newProduct.Description">

              <label>Price:</label>
              <input type="number" [(ngModel)]="newProduct.Price">

              <label>Quantity:</label>
              <input type="number" [(ngModel)]="newProduct.Quantity">

              <div class="modal-actions">
                  <button class="primary" (click)="addProduct()">Save</button>
                  <button class="secondary" (click)="toggleModal()">Cancel</button>
              </div>
          </div>
      </div>
  </div>

  <!-- Product List -->
  <section class="results">
      <app-product 
          *ngFor="let product of products" 
          [product]="product" 
          (click)="viewProduct(product)">
      </app-product>
  </section>

  <!-- Product Details Modal -->
  <div class="modal-overlay" *ngIf="isProductModalOpen" (click)="closeProductModal()">
      <div class="modal" (click)="$event.stopPropagation()">
          <div class="modal-content">
              <h2>Product Details</h2>
              <p><strong>Name:</strong> {{ selectedProduct?.Name }}</p>
              <p><strong>Description:</strong> {{ selectedProduct?.Description }}</p>
              <p><strong>Price:</strong> {{ selectedProduct?.Price}} /-</p>
              <p><strong>Quantity:</strong> {{ selectedProduct?.Quantity }}</p>

              <div class="modal-actions">
                  <button class="secondary" (click)="closeProductModal()">Close</button>
              </div>
          </div>
      </div>
  </div>
  `,
  styleUrls: ['./home.component.css']
})
export class HomeComponent { 

  products: Product[] = [];
  excelService: ExcelService = inject(ExcelService);
  router: Router = inject(Router);
  
  generateUUID(): string {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
      const r = Math.random() * 16 | 0, v = c === 'x' ? r : (r & 0x3 | 0x8);
      return v.toString(16);
    });
  }

  isModalOpen = false;
  isProductModalOpen = false;
  selectedProduct: Product | null = null;
  
  newProduct: Product = {ID: this.generateUUID(), Name: '', Price: 0, Quantity: 0, Description: "" };
  
  constructor() {
    this.loadProducts();
  }
  
  navigateToReport() {
    this.router.navigate(['/report']);
  }

  async loadProducts() {
    await this.excelService.fetchProducts();
    this.excelService.products$.subscribe((data: Product[]) => {
      this.products = data;
    });
  }

  toggleModal(): void {
    this.isModalOpen = !this.isModalOpen;
  }

  async addProduct() {
    if (!this.newProduct.Name) {
      alert("Please enter a product name!");
      return;
    }
    if (this.newProduct.Price <= 0) {
      alert("Please enter a valid price greater than 0!");
      return;
    }
    if (this.newProduct.Quantity < 0) {
      alert("Please enter a valid quantity!");
      return;
    }
    
    await this.excelService.addProduct([this.newProduct]);
    
    this.toggleModal();
    this.loadProducts();
    window.location.reload()
    this.newProduct = {ID: this.generateUUID(), Name: '', Price: 0, Quantity: 0, Description: "" };
  }

  // Opens product details modal
  viewProduct(product: Product) {
    this.selectedProduct = product;
    this.isProductModalOpen = true;
  }

  // Closes product details modal
  closeProductModal() {
    this.isProductModalOpen = false;
    this.selectedProduct = null;
  }
}
