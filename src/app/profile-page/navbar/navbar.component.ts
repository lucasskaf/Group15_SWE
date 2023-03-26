import { Component, Output, Input, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  @Output() openSidenav = new EventEmitter<boolean>();
  // @Input() sidenavStatus: boolean = false;
  @Input() sidenavStatus: boolean = true;

  constructor() { }

  ngOnInit() {

  }

  public onSidenavClick(): void {
    this.openSidenav.emit(!this.sidenavStatus);
    this.sidenavStatus = !this.sidenavStatus
  }
}