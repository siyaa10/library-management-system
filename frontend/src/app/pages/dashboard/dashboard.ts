
import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { forkJoin } from 'rxjs';
import { LibraryService } from '../../services/library.service';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, RouterLink],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.scss'
})
export class Dashboard implements OnInit {

  stats = [
    { title: 'Total Books', value: 0, icon: '📚', color: 'blue' },
    { title: 'Total Members', value: 0, icon: '👥', color: 'purple' },
    { title: 'Books Issued', value: 0, icon: '📖', color: 'orange' },
    { title: 'Overdue Books', value: 0, icon: '⏰', color: 'red' }
  ];

  activities: any[] = [];
  error = '';
  today = new Date();

  availabilityChart: {
    label: string;
    count: number;
    percentage: number;
    color: string;
  }[] = [];

  borrowingChart: {
    label: string;
    count: number;
    percentage: number;
    color: string;
  }[] = [];

  categoryChart: {
    name: string;
    count: number;
    percentage: number;
  }[] = [];

  constructor(
    private library: LibraryService,
    private cdr: ChangeDetectorRef
  ) {}

  ngOnInit(): void {
    this.loadDashboard();
  }

  loadDashboard(): void {
    forkJoin({
      dashboard: this.library.getDashboard(),
      books: this.library.getBooks(),
      members: this.library.getMembers(),
      loans: this.library.getLoans()
    }).subscribe({
      next: data => {

        // Dashboard statistics
        this.stats[0].value = Number(data.dashboard.totalBooks);
        this.stats[1].value = Number(data.dashboard.totalMembers);
        this.stats[2].value = Number(data.dashboard.issuedBooks);
        this.stats[3].value = Number(data.dashboard.overdueBooks);

        // Recent activity
        this.activities = data.loans.slice(-5).reverse().map(loan => ({
          title: data.books.find(book => book.id === loan.bookId)?.title || `Book #${loan.bookId}`,
          member: data.members.find(member => member.id === loan.memberId)?.name || `Member #${loan.memberId}`,
          action: loan.status === 'Returned' ? 'Book returned' : 'Book issued',
          date: loan.returnedAt || loan.issueDate,
          status: loan.status
        }));

        // Prepare all charts
        this.prepareCharts(data.books, data.loans);

        this.error = '';
        this.cdr.markForCheck();
      },
      error: err => {
        console.error('Dashboard loading error:', err);
        this.error = 'Unable to load dashboard data. Please check the backend connection.';
        this.cdr.markForCheck();
      }
    });
  }

  prepareCharts(books: any[], loans: any[]): void {

    // Book availability
    const totalCopies = books.reduce(
      (sum, book) => sum + Number(book.totalCopies || 0), 0
    );

    const availableCopies = books.reduce(
      (sum, book) => sum + Number(book.availableCopies || 0), 0
    );

    const issuedCopies = totalCopies - availableCopies;

    this.availabilityChart = [
      {
        label: 'Available Copies',
        count: availableCopies,
        percentage: totalCopies ? availableCopies / totalCopies * 100 : 0,
        color: 'available'
      },
      {
        label: 'Issued Copies',
        count: issuedCopies,
        percentage: totalCopies ? issuedCopies / totalCopies * 100 : 0,
        color: 'issued'
      }
    ];

    // Books by category
    const categoryCounts: { [key: string]: number } = {};

    books.forEach(book => {
      const category = book.category || 'Other';
      categoryCounts[category] = (categoryCounts[category] || 0) + 1;
    });

    const totalBooks = books.length;

    this.categoryChart = Object.entries(categoryCounts)
      .map(([name, count]) => ({
        name,
        count,
        percentage: totalBooks ? Math.round(count / totalBooks * 100) : 0
      }))
      .sort((a, b) => b.count - a.count);

    // Borrowing status
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    const returned = loans.filter(
      loan => loan.status === 'Returned'
    ).length;

    const overdue = loans.filter(loan => {
      if (loan.status !== 'Issued') return false;

      const dueDate = new Date(loan.dueDate);
      dueDate.setHours(0, 0, 0, 0);

      return dueDate < today;
    }).length;

    const active = loans.filter(loan => {
      if (loan.status !== 'Issued') return false;

      const dueDate = new Date(loan.dueDate);
      dueDate.setHours(0, 0, 0, 0);

      return dueDate >= today;
    }).length;

    const totalLoans = active + returned + overdue;

    this.borrowingChart = [
      {
        label: 'Active',
        count: active,
        percentage: totalLoans ? active / totalLoans * 100 : 0,
        color: 'active'
      },
      {
        label: 'Returned',
        count: returned,
        percentage: totalLoans ? returned / totalLoans * 100 : 0,
        color: 'returned'
      },
      {
        label: 'Overdue',
        count: overdue,
        percentage: totalLoans ? overdue / totalLoans * 100 : 0,
        color: 'overdue'
      }
    ];
  }
}
