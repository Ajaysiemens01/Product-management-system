import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { Product } from './product';
import { map, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ExcelService {
  private apiUrl = 'http://localhost:8081/api/products';

  private httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/vnd.api+json',
      'X-API-KEY': 'product-management-api-key'
    })
  };

  private productsSubject = new BehaviorSubject<Product[]>([]);
  products$ = this.productsSubject.asObservable();

  constructor(private http: HttpClient) {
    console.log(" ExcelService Initialized");
  }

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
        tap(products => console.log(" Processed Products:", products))
      )
      .subscribe({
        next: (products) => this.productsSubject.next(products),
        error: (error) => console.error(" Error fetching products:", error)
      });
  }
}
