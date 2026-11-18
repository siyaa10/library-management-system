import { CommonModule } from '@angular/common';
import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { LibraryService } from '../../services/library.service';
import { Book } from '../../models/book';
import { Loan } from '../../models/loan';

@Component({
  selector: 'app-book-details',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './book-details.html',
  styleUrl: './book-details.scss'
})
export class BookDetails implements OnInit {
  book: Book | null = null;
  history: Loan[] = [];
  loading = true;

  constructor(
  private route: ActivatedRoute,
  private router: Router,
  private libraryService: LibraryService,
  private cdr: ChangeDetectorRef
) {}

  ngOnInit(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));

    if (!id) {
      this.router.navigate(['/books']);
      return;
    }

    this.libraryService.getBookById(id).subscribe({
      next: book => {
  this.book = book;
  this.loading = false;
  this.cdr.markForCheck();
},
     error: () => {
  this.loading = false;
  this.cdr.markForCheck();
}
    });

   this.libraryService.getBookHistory(id).subscribe({
  next: history => {
    this.history = history;
    this.cdr.markForCheck();
  },
  error: () => {
    this.history = [];
    this.cdr.markForCheck();
  }
});
  }

  goBack(): void {
    this.router.navigate(['/books']);
  }
}