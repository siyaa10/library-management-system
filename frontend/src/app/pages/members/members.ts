import { Router } from '@angular/router';
import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Member } from '../../models/member';
import { LibraryService } from '../../services/library.service';

@Component({
  selector: 'app-members',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './members.html',
  styleUrl: './members.scss'
})
export class Members implements OnInit {
  members: Member[] = [];
  searchText = '';
  showForm = false;
  editingId: number | null = null;
  message = '';
  error = '';

  form: Member = this.emptyForm();

  constructor(
  private library: LibraryService,
  private cdr: ChangeDetectorRef,
  private router: Router
) {}

  ngOnInit(): void {
    this.loadMembers();
  }

  emptyForm(): Member {
    return {
      name: '',
      email: '',
      phone: '',
      address: '',
      membershipType: 'Regular'
    };
  }

  loadMembers(): void {
  this.library.getMembers().subscribe({
    next: data => {
      this.members = data;
      this.cdr.markForCheck();
    },
    error: () => {
      this.error = 'Unable to load members.';
      this.cdr.markForCheck();
    }
  });
}

  get filteredMembers(): Member[] {
    const term = this.searchText.toLowerCase();

    return this.members.filter(member =>
      member.name.toLowerCase().includes(term) ||
      (member.email || '').toLowerCase().includes(term) ||
      member.phone.includes(term)
    );
  }

  openAddForm(): void {
    this.form = this.emptyForm();
    this.editingId = null;
    this.showForm = true;
    this.message = '';
    this.error = '';
  }

  openEditForm(member: Member): void {
    this.form = { ...member };
    this.editingId = member.id!;
    this.showForm = true;
    this.message = '';
    this.error = '';
  }

  saveMember(): void {
    if (!this.form.name.trim() || !this.form.phone.trim()) {
      this.error = 'Name and phone number are required.';
      return;
    }

    const request = this.editingId
      ? this.library.updateMember(this.editingId, this.form)
      : this.library.addMember(this.form);

    request.subscribe({
      next: () => {
        this.message = this.editingId
          ? 'Member updated successfully!'
          : 'Member registered successfully!';
        this.showForm = false;
        this.loadMembers();
      },
      error: err => {
        this.error = err.error?.error || 'Unable to save member.';
      }
    });
  }

  deleteMember(id: number): void {
    if (!confirm('Are you sure you want to delete this member?')) return;

    this.library.deleteMember(id).subscribe({
      next: () => {
        this.message = 'Member deleted successfully!';
        this.loadMembers();
      },
      error: err => {
  this.error = err.error?.error || 'Unable to delete member.';
  this.message = '';
  this.cdr.markForCheck();
}
    });
  }
  viewMember(id: number): void {
  this.router.navigate(['/members', id]);
}
}
