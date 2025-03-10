import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, Observable, ReplaySubject } from 'rxjs';
import { Product } from './product';
import { map, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ExcelService {
  private apiUrl = 'http://localhost:8081/api/products';
  private updateApiUrl = 'http://localhost:8082/api/inventory';
  private deductApiUrl = 'http://localhost:8084/api/purchase';
  private reportApiUrl = 'http://localhost:8083/api/report/inventory';
  private httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/vnd.api+json',
      'X-API-KEY': 'product-management-api-key'
    })
  };

  private productsSubject = new ReplaySubject<Product[]>(1);
  products$ = this.productsSubject.asObservable();
  private reportSubject = new ReplaySubject<Product[]>(1);
  report$ = this.reportSubject.asObservable();
  constructor(private http: HttpClient) {
    console.log(" ExcelService Initialized");
  }
  //Fetch Products form url
  fetchProducts(): void {
    this.http.get<{ data: any[] }>(this.apiUrl, this.httpOptions)
      .pipe(
        map(response => response.data?.map(item => ({
          ID: item.id,
          Name: item.attributes.name,
          Description: item.attributes.description,
          Price: item.attributes.price,
          Quantity: item.attributes.quantity
        })) || []),
        tap(products => console.log(products))
      )
      .subscribe({
        next: (products) => {
          this.productsSubject.next(products);
        },
        error: (error) => console.error("Error fetching products:", error)
      });
  }

  // Add products received from URL
  addProduct(products: Product[]): void {
    const requestBody = {
      data: products.map(p => ({
        type: 'product',
        attributes: {
          name: p.Name,
          description: p.Description || "",
          price: p.Price,
          quantity: p.Quantity
        }
      }))
    };

    this.http.post(this.apiUrl, requestBody, this.httpOptions)
      .subscribe({
        next: (response) => {
          console.log("Products added successfully:", response);
        },
        error: (error) => {
          console.error("Error adding products:", error);
        }
      });
  }
  // Update product by ID
  updateProduct(id: string, updatedProduct: Product):void {
    const requestBody = {
      data: {
        type: 'product',
        attributes: {
          name: updatedProduct.Name,
          description: updatedProduct.Description || "",
          price: updatedProduct.Price,
          quantity: updatedProduct.Quantity
        }
      }
    };
   
     this.http.put(`${this.updateApiUrl}/${id}`, requestBody, this.httpOptions)
     .subscribe({
       next: (response) => {
         console.log("Products updated successfully:", response);
       },
       error: (error) => {
         console.error("Error updating products:", error);
       }
     });

  }

  //Deduct Product Quantity
 deductProduct(id: string, changeVal: number):void {
    const requestBody = {
      data: {
        type: 'product-change',
        attributes: {
          change:changeVal
        }
      }
    };
   
     this.http.put(`${this.deductApiUrl}/${id}`, requestBody, this.httpOptions)
     .subscribe({
       next: (response) => {
         console.log("Quantity successfully:", response);
       },
       error: (error) => {
         console.error("Error deducting products:", error);
       }
     });

  }
  // Fetch products with low stock
 fetchLowStockProducts(threshold: number): void {
  this.http.get<{ data: any[] }>(`${this.reportApiUrl}?restock_threshold=${threshold}`, this.httpOptions)
    .pipe(
      map(response => response.data?.map(item => ({
        ID: item.id,
        Name: item.attributes.name,
        Description: item.attributes.description,
        Price: item.attributes.price,
        Quantity: item.attributes.quantity
      })) || []),
      tap(report => console.log("Low stock products:",report))
    )
    .subscribe({
      next: (report) => {
        this.reportSubject.next(report);
      },
      error: (error) => console.error("Error fetching low stock products:", error)
    });
}


}

