import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';
import { LibraryService } from '../../services/library.service';
import { Member } from '../../models/member';
import { Loan } from '../../models/loan';

@Component({
  selector: 'app-member-details',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './member-details.html',
  styleUrl: './member-details.scss'
})
export class MemberDetails implements OnInit {
  member: Member | null = null;
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
      this.router.navigate(['/members']);
      return;
    }

    this.libraryService.getMemberById(id).subscribe({
      next: member => {
        this.member = member;
        this.loading = false;
        this.cdr.markForCheck();
      },
      error: () => {
        this.loading = false;
        this.cdr.markForCheck();
      }
    });

    this.libraryService.getMemberHistory(id).subscribe({
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
    this.router.navigate(['/members']);
  }
}