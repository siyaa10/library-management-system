
import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Loan } from '../../models/loan';
import { LibraryService } from '../../services/library.service';

@Component({
  selector: 'app-history',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './history.html',
  styleUrl: './history.scss'
})
export class History implements OnInit {
  loans: Loan[] = [];
  searchText = '';
  selectedStatus = 'All';
  error = '';

  constructor(
  private library: LibraryService,
  private cdr: ChangeDetectorRef
) {}

  ngOnInit(): void {
  this.library.getLoans().subscribe({
    next: data => {
      this.loans = data;
      this.cdr.markForCheck();
    },
    error: () => {
      this.error = 'Unable to load borrowing history.';
      this.cdr.markForCheck();
    }
  });
}

  get filteredLoans(): Loan[] {
    return this.loans.filter(loan => {
      const matchesSearch =
        String(loan.id).includes(this.searchText) ||
        String(loan.bookId).includes(this.searchText) ||
        String(loan.memberId).includes(this.searchText);

      const matchesStatus =
        this.selectedStatus === 'All' ||
        loan.status === this.selectedStatus;

      return matchesSearch && matchesStatus;
    });
  }
}
