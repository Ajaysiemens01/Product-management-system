import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Product } from '../product';
@Component({
  selector: 'app-product',
  imports: [CommonModule],
  standalone: true,
  template: `
  <section class="listing">
    <h2 class="listing-heading">{{ product.Name }}</h2>
    <section class = "listing-content">
    <p class="listing-data"> Price: {{ product.Price}}/-</p>
    <p class="listing-data"> Quantity: {{product.Quantity}}</p>
    </section>
    
    <section class="content">
      <button class="primary" type="button">Edit</button>
      <br/>
      <button class="primary" type="button">Deduct</button>
  </section>
  </section>
  `,
  styleUrl: './product.component.css',
})
export class ProductComponent {
  @Input() product!: Product;
}
