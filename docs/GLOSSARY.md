# Glossary for Catalog Service

This glossary defines key terms and concepts used within the **Catalog Service** to help developers and stakeholders understand essential terminology.


---

### Categories
A logical grouping of **menu items** used to organize items in a menu (e.g., Appetizers, Main Course, Beverages). Each category can be specific to a **restaurant** or **country**, allowing for localized menus.

### Menu Items
The core products or dishes offered by a restaurant. A **menu item** includes attributes like name, description, preparation time, and availability status. Menu items can belong to a category and may have associated prices, modifiers, and allergens.

### Item Prices
The pricing information for a **menu item**. **Item prices** can vary by **restaurant** and **country** and may also include time-based pricing (e.g., for happy hours or seasonal items). Each item price entry supports multiple currencies and defines a start and optional end date for pricing.

### Modifiers
Customizable options available for **menu items** (e.g., extra cheese, add sauce, larger portion size). Modifiers allow customers to customize items based on their preferences and can have additional costs associated with them.

### Modifier Prices
Pricing information for specific **modifiers**. This allows restaurants to apply an additional charge for certain modifications (e.g., extra toppings). Modifier prices support time-based configurations and can vary by restaurant and country.

### Promotions
Discounts or special offers applied to **menu items** or **categories**. Promotions can be based on a percentage discount or a fixed amount and are typically time-limited. Promotions allow restaurants to create incentives for customers and can be specific to a restaurant or country.

### Menus
A collection of **menu items** tailored to a specific time or purpose, such as breakfast, lunch, dinner, or a seasonal event. Menus help organize items available during certain times or events, improving customer experience.

### Allergens
Information about potential allergens present in a **menu item** (e.g., peanuts, dairy, gluten). **Allergens** help customers with dietary restrictions make informed choices and enhance food safety.

### Dietary Flags
Labels that describe dietary restrictions or preferences for **menu items** (e.g., vegan, gluten-free, keto-friendly). Dietary flags help customers easily identify items that meet their dietary needs.

### Availability
Defines whether a **menu item** is available for order. **Availability** can be specific to certain dates or times, allowing for seasonal or limited-time items. This feature helps manage items that are not consistently offered.

### Item Modifiers
A many-to-many relationship between **menu items** and **modifiers**. It links specific modifiers to items, allowing each item to have its own set of customizable options.

### Menu Item-Menu Association
The relationship that links **menu items** to specific **menus**. This association allows an item to appear in one or more menus (e.g., an item available for both breakfast and lunch).

### Category Modifiers
Modifiers that apply to all **menu items** within a specific **category**. This association simplifies applying common modifiers (e.g., add cheese to all sandwiches in the Sandwich category) without having to set them individually for each item.

### Restaurant
A distinct business entity within the system that can have its own **menu items**, **categories**, **modifiers**, and **promotions**. Each **restaurant** can customize its catalog independently.

### Country Code
A standard two-letter code (ISO 3166-1 alpha-2) representing the country. **Country codes** help localize **menu items**, prices, and promotions, allowing for unique configurations based on regional requirements.

### Currency Code
A three-letter code (ISO 4217) representing the currency used for pricing **menu items** and **modifiers** (e.g., USD, EUR, IDR). **Currency codes** allow pricing information to be displayed and charged in multiple currencies as needed for different countries.

### UUID
A universally unique identifier used as the primary key for each table entry. **UUIDs** ensure uniqueness across distributed systems, supporting scalability and helping to avoid conflicts between records across different databases.

---

This glossary provides clear definitions for terms and concepts essential to the **Catalog Service** and its functionalities. It serves as a reference for developers, stakeholders, and anyone interacting with the system.

---
For technical details and usage, refer to `ARCHITECTURE.md` and `API_DOCS.md`.
