import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {

  @Output() isClosed = new EventEmitter<boolean>();
  // @Input() isOpen : boolean = false
  @Input() isOpen : boolean = true

  constructor() { }

  ngOnInit(): void {

  }

  toggleSidenav() {
    this.isOpen = !this.isOpen
    this.isClosed.emit(this.isOpen)
  }
}