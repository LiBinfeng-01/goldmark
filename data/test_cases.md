# Complex Markdown Test Cases

## 1. Nested Headings and Structure

### Level 3 Heading
This is a paragraph under a level 3 heading. It contains some basic text that should be chunked properly.

#### Level 4 Heading
This is a paragraph under a level 4 heading. It demonstrates how the chunker handles deeper nesting levels.

##### Level 5 Heading
This is a paragraph under a level 5 heading. The chunker should maintain the hierarchical structure even at this level.

## 2. Code Blocks and Technical Content

Here's a code block with some Python code:

```python
def process_data(data):
    """
    This is a complex function that processes data.
    It has multiple lines and should be kept together.
    """
    result = []
    for item in data:
        if item > 0:
            result.append(item * 2)
    return result
```

And here's some JavaScript:

```javascript
const processData = (data) => {
    // This is a complex function
    return data
        .filter(x => x > 0)
        .map(x => x * 2)
        .reduce((a, b) => a + b, 0);
};
```

## 3. Lists and Nested Content

### Ordered List
1. First item with some detailed explanation that might need to be chunked properly
2. Second item with a sublist:
   - Subitem 1
   - Subitem 2
   - Subitem 3
3. Third item with a code block:
   ```bash
   echo "Hello World"
   ```

### Unordered List
* Main point 1
  * Subpoint 1.1
  * Subpoint 1.2
* Main point 2
  * Subpoint 2.1
  * Subpoint 2.2

## 4. Tables and Structured Data

| Header 1 | Header 2 | Header 3 |
|----------|----------|----------|
| Cell 1   | Cell 2   | Cell 3   |
| Cell 4   | Cell 5   | Cell 6   |

## 5. Long Paragraphs

This is a very long paragraph that should be split into multiple chunks based on the token limit. It contains a lot of text that needs to be processed properly. The chunker should maintain the semantic meaning while splitting this content. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## 6. Special Characters and Formatting

### Text Formatting
*This is italic text* and **this is bold text**. We can also have ***bold and italic*** text.

### Links and References
[Link to Google](https://www.google.com)

### Images
![Alt text](https://example.com/image.jpg)

## 7. Mixed Content

### Section with Multiple Elements
This section contains a mix of different elements:

1. A numbered list item
2. A paragraph with *italic* and **bold** text
3. A code block:
   ```python
   print("Hello, World!")
   ```
4. A table:
   | A | B |
   |---|---|
   | 1 | 2 |

### Complex Nested Structure
- Level 1
  - Level 2
    - Level 3
      - Level 4
        - Level 5
          This is a very deep nested structure that should be handled properly by the chunker.

## 8. Edge Cases

### Empty Sections
#### Empty Heading

### Single Character Sections
#### A
This is a section with a single character heading.

### Very Long Words
This is a section with a very long word: pneumonoultramicroscopicsilicovolcanoconiosis

### Special Characters
This section contains special characters: !@#$%^&*()_+-=[]{}|;:,.<>?

### Unicode Characters
This section contains Unicode characters: 你好世界, こんにちは, 안녕하세요 