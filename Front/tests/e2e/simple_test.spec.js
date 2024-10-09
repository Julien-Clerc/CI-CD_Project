import { test, expect } from '@playwright/test';

test.describe('Simple E2E Test', () => {
    test('check title', async ({ page }) => {
        await page.goto('http://localhost:5173');

        await expect(page.locator('text=Accueil')).toBeVisible()
    })
});

test('should interface with the backend', async ({ request}) => {
    const response = await request.get('http://localhost:3000/config');

    expect(response.ok()).toBeTruthy();
    const responseBody = await response.json();
})
