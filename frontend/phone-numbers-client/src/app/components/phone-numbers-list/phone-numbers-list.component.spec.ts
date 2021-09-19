import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PhoneNumbersListComponent } from './phone-numbers-list.component';

describe('PhoneNumbersListComponent', () => {
  let component: PhoneNumbersListComponent;
  let fixture: ComponentFixture<PhoneNumbersListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PhoneNumbersListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PhoneNumbersListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
