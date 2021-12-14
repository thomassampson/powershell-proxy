// okta-signin.spec.ts created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test

describe("Okta Sign In", () => {
  it("Load Okta Device SignIn", () => {
    cy.visit("https://one.acuityads.cloud/activate")
      .get("#user-code")
      .type("JGHRMGGN")
      .get(".button")
      .click()
      .get("#okta-signin-username")
      .type("thomas.sampson")
      .get("#okta-signin-password")
      .type("CaUkGuy@2807=")
      .get("#okta-signin-submit")
      .click();
  });
});
