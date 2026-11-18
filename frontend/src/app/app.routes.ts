
import { Routes } from '@angular/router';
import { Dashboard } from './pages/dashboard/dashboard';
import { Books } from './pages/books/books';
import { Members } from './pages/members/members';
import { Loans } from './pages/loans/loans';
import { History } from './pages/history/history';
import { BookDetails } from './pages/book-details/book-details';
import { MemberDetails } from './pages/member-details/member-details';

export const routes: Routes = [
  { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
  { path: 'dashboard', component: Dashboard },
  { path: 'books', component: Books },
  { path: 'books/:id', component: BookDetails },
  { path: 'members', component: Members },
  { path: 'members/:id', component: MemberDetails },
  { path: 'loans', component: Loans },
  { path: 'history', component: History },
  { path: '**', redirectTo: 'dashboard' }
];
