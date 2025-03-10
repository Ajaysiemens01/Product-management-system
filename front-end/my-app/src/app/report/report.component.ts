import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ExcelService } from '../excel.service';
import { Product } from '../product';
import { ProductComponent } from '../product/product.component';
import * as ExcelJS from 'exceljs';
import { saveAs } from 'file-saver';

@Component({
  selector: 'app-report',
  standalone: true,
  imports: [CommonModule, FormsModule, ProductComponent],
  template: `
    <section class="content">
      <h2>Generate Low Stock Report</h2>

      <label class="listing-data">Enter Minimum Stock:</label>
      <input type="number" [(ngModel)]="threshold" min="1"><br>
      <section class="button-content">
      <button class="primary" (click)="generateReport()">Generate Report</button>
      </section>  
    </section>

    <section class="results" *ngIf="lowStockProducts.length > 0">
      <app-product *ngFor="let product of lowStockProducts" [product]="product" [isreport]="true"></app-product>
</section>
<section class="content">
      <section class="button-content" *ngIf="lowStockProducts.length > 0">
      <button class="primary" (click)="downloadReport()" >Download Report</button>
      </section>  
    </section>

    <p *ngIf="lowStockProducts.length === 0 && reportGenerated">No products below the threshold.</p>
  `,
  styleUrls: ['./report.component.css'],
})
export class ReportComponent {
    
    excelService: ExcelService = inject(ExcelService);
    threshold: number = 20;
    lowStockProducts: Product[] = [];
    reportGenerated = false;

    constructor() {}

    async generateReport() {
      if (this.threshold <= 0) {
        alert('Please enter a valid threshold greater than 0.');
        return;
      }
      
      await this.excelService.fetchLowStockProducts(this.threshold);
      this.excelService.report$.subscribe((data: Product[]) => {
        this.lowStockProducts = data;
        this.reportGenerated = true;
      });
    }

     async downloadReport() {
      if (this.threshold <= 0) {
        alert('Please enter a valid threshold greater than 0.');
        return;
      }
      const workbook = new ExcelJS.Workbook();
      const worksheet = workbook.addWorksheet('Low Stock Report');

      // Add Headers
      worksheet.addRow(['ID', 'Name', 'Description', 'Price', 'Quantity']);

      // Add Data
      this.lowStockProducts.forEach(product => {
        worksheet.addRow([
          product.ID,
          product.Name,
          product.Description,
          product.Price,
          product.Quantity
        ]);
      });

      // Style Headers
      worksheet.getRow(1).font = { bold: true };

      // Generate Excel File
      const buffer = await workbook.xlsx.writeBuffer();
      const fileName = `Low_Stock_Report.xlsx`;

      saveAs(new Blob([buffer]), fileName);
      
    }
}
