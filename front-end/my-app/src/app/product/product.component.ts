import { Component, Input, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Product } from '../product';
import { ExcelService } from '../excel.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-product',
  imports: [CommonModule, FormsModule],
  standalone: true,
  template: `
  <!-- Click anywhere on product to open details modal -->
  <section class="listing" (click)="openDetailsModal($event)">
    <h2 class="listing-heading">{{ product.Name }}</h2>
    <section class="listing-content">
      <p class="listing-data"> Price : {{ product.Price }} /-</p>
      <p class="listing-data"> Quantity : {{ product.Quantity }}</p>
    </section>

    <section class="content" *ngIf="!isreport">
      <button class="primary" type="button" (click)="openEditModal($event)">Edit</button>
      <br/>
      <button class="primary" type="button" (click)="openDeductModal($event)">Deduct</button>
    </section>
  </section>

  <!-- Product Details Modal -->
  <div class="modal-overlay" *ngIf="isDetailsModalOpen" (click)="closeDetailsModal($event)">
    <div class="modal" (click)="$event.stopPropagation()">
      <div class="modal-content">
        <h2>Product Details</h2>
        <p><strong>Name:</strong> {{ product.Name }}</p>
        <p><strong>Description:</strong> {{ product.Description }}</p>
        <p><strong>Price:</strong> {{ product.Price }} /-</p>
        <p><strong>Quantity:</strong> {{ product.Quantity }}</p>

        <div class="modal-actions">
          <button class="secondary" (click)="closeDetailsModal($event)">Close</button>
        </div>
      </div>
    </div>
  </div>

  <!-- Edit Modal -->
  <div class="modal-overlay" *ngIf="isEditModalOpen" (click)="closeEditModal($event)">
    <div class="modal" (click)="$event.stopPropagation()">
      <div class="modal-content">
        <h2>Edit Product</h2>
        <label>Name:</label>
        <input type="text" [(ngModel)]="editedProduct.Name">
        <label>Description:</label>
        <input type="text" [(ngModel)]="editedProduct.Description">
        <label>Price:</label>
        <input type="number" [(ngModel)]="editedProduct.Price">
        <label>Quantity:</label>
        <input type="number" [(ngModel)]="editedProduct.Quantity">

        <div class="modal-actions">
          <button class="primary" (click)="updateProduct()">Save</button>
          <button class="secondary" (click)="closeEditModal($event)">Cancel</button>
        </div>
      </div>
    </div>
  </div>

  <!-- Deduct Modal -->
  <div class="modal-overlay" *ngIf="isDeductModalOpen" (click)="closeDeductModal($event)">
    <div class="modal" (click)="$event.stopPropagation()">
      <div class="modal-content">
        <h2>Deduct</h2>
        <label>Deduct Quantity:</label>
        <input type="number" [(ngModel)]="change">
        
        <div class="modal-actions">
          <button class="primary" (click)="deductProduct()">Save</button>
          <button class="secondary" (click)="closeDeductModal($event)">Cancel</button>
        </div>
      </div>
    </div>
  </div>
  `,
  styleUrl: './product.component.css',
})
export class ProductComponent {
  @Input() product!: Product;
  @Input() isreport: boolean = false;

  isEditModalOpen = false;
  isDeductModalOpen = false;
  isDetailsModalOpen = false;  // New state for product details modal
  editedProduct!: Product;
  change!: number;

  excelService: ExcelService = inject(ExcelService);

  // Open product details modal (Click anywhere on the product)
  openDetailsModal(event: Event) {
    event.stopPropagation(); // Prevent bubbling
    this.isDetailsModalOpen = true;
  }

  // Close product details modal
  closeDetailsModal(event: Event) {
    event.stopPropagation();
    this.isDetailsModalOpen = false;
  }

  // Open edit modal and prevent bubbling
  openEditModal(event: Event) {
    event.stopPropagation(); 
    this.isEditModalOpen = true;
    this.editedProduct = { ...this.product };
  }

  // Close edit modal
  closeEditModal(event: Event) {
    event.stopPropagation();
    this.isEditModalOpen = false;
  }

  // Open deduct modal
  openDeductModal(event: Event) {
    event.stopPropagation();
    this.isDeductModalOpen = true;
    this.change = 0;
  }

  // Close deduct modal
  closeDeductModal(event: Event) {
    event.stopPropagation();
    this.isDeductModalOpen = false;
  }

  // Update product
  async updateProduct() {
    if (!this.editedProduct.Name || this.editedProduct.Price <= 0 || this.editedProduct.Quantity < 0) {
      alert("Please enter valid product details!");
      return;
    }
    
    
  await this.excelService.updateProduct(this.product.ID, this.editedProduct);
  window.location.reload();
  }

  // Deduct product quantity
  async deductProduct() {
    if (this.change > this.product.Quantity) {
      alert("You are out of stock");
      return;
    }

    await this.excelService.deductProduct(this.product.ID, this.change);
    window.location.reload();
  }
}
