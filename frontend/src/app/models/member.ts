
export interface Member {
  id?: number;
  name: string;
  email?: string | null;
  phone: string;
  address: string;
  membershipType: string;
  membershipDate?: string;
  membershipStatus?: string;
}
