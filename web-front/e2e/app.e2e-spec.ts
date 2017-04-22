import { WebFrontPage } from './app.po';

describe('web-front App', () => {
  let page: WebFrontPage;

  beforeEach(() => {
    page = new WebFrontPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
