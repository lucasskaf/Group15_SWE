import { Component, OnInit, EventEmitter, Output, Input} from '@angular/core';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  @Output() openSidenav = new EventEmitter<boolean>();
  @Input() sidenavStatus: boolean = false

  constructor() { }

  ngOnInit() {

  }

  public onSidenavClick(): void {
    this.openSidenav.emit(!this.sidenavStatus);
    this.sidenavStatus = !this.sidenavStatus
  }
}
