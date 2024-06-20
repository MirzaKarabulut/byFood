import React, { useState, useEffect } from "react";
import axios from "axios";
import CardComponent from "./CardComponent";

interface Book {
  ID: number;
  title: string;
  author: string;
  description: string;
  releaseDate: string;
}

interface UserInterfaceProps {
  backendName: string;
}

const UserInterface: React.FC<UserInterfaceProps> = ({ backendName }) => {
  const apiURL = "http://localhost:8080";
  const [books, setBooks] = useState<Book[]>([]);
  const [newBook, setNewBook] = useState({ title: "", description: "", author: "", releaseDate: "" });
  const [updateBook, setUpdateBook] = useState({ id: "", title: "", description: "", author: "", releaseDate: "" });
  const [error, setError] = useState<string | null>(null);

  // Define styles based on the backendName
  const backgroundColors: { [key: string]: string } = {
    go: 'bg-cyan-500',
  };

  const buttonColors: { [key: string]: string } = {
    go: 'bg-cyan-700 hover:bg-blue-600',
  };

  const bgColor = backgroundColors[backendName as keyof typeof backgroundColors] || 'bg-gray-200';
  const btnColor = buttonColors[backendName as keyof typeof buttonColors] || 'bg-gray-700 hover:bg-gray-600';

  // Fetch all books
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(`${apiURL}/books`);
        setBooks(response.data.reverse());
      } catch (error) {
        setError("Error fetching data");
        console.error("Error fetching data", error);
      }
    };
    fetchData();
  }, [backendName, apiURL]);

  // Validate book input
  const validateBook = (book: any) => {
    if (!book.title || !book.author || !book.description || !book.releaseDate) {
      setError("All fields are required");
      return false;
    }
    setError(null);
    return true;
  };

  // Create a new book
  const createBook = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!validateBook(newBook)) return;

    try {
      const response = await axios.post(`${apiURL}/books`, newBook);
      setBooks([response.data, ...books]);
      setNewBook({ title: "", description: "", author: "", releaseDate: "" });
    } catch (error) {
      setError("Error creating book");
      console.error("Error creating book:", error);
    }
  };

  // Update a book
  const handleUpdateBook = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!validateBook(updateBook)) return;

    try {
      await axios.put(`${apiURL}/books/${updateBook.id}`, { title: updateBook.title, author: updateBook.author, description: updateBook.description, releaseDate: updateBook.releaseDate });
      setUpdateBook({ id: "", title: "", description: "", author: "", releaseDate: "" });
      setBooks(
        books.map((book) => {
          if (book.ID === parseInt(updateBook.id)) {
            return { ...book, title: updateBook.title, author: updateBook.author, description: updateBook.description, releaseDate: updateBook.releaseDate };
          }
          return book;
        })
      );
    } catch (error) {
      setError("Error updating book");
      console.error('Error updating book:', error);
    }
  };

  // Delete a book
  const deleteBook = async (bookID: number) => {
    try {
      await axios.delete(`${apiURL}/books/${bookID}`);
      setBooks(books.filter((book) => book.ID !== bookID));
    } catch (error) {
      setError("Error deleting book");
      console.error('Error deleting book:', error);
    }
  };

  return (
    <div className={`user-interface ${bgColor} ${backendName} w-full max-w-md p-4 my-4 rounded shadow`}>
      <h2 className="text-xl font-bold text-center text-white mb-6">{`${backendName.charAt(0).toUpperCase() + backendName.slice(1)} Backend`}</h2>

      {error && <div className="mb-4 p-2 text-red-600 bg-red-200 rounded">{error}</div>}

      {/* Get all books */}
      <button onClick={() => window.location.reload()} className={`${btnColor} text-white py-2 px-4 rounded`}>
        Refresh Books
      </button>

      {/* Create a new book */}
      <form onSubmit={createBook} className="mb-6 p-4 bg-blue-100 rounded shadow">
        <input
          placeholder="Title"
          value={newBook.title}
          onChange={(e) => setNewBook({ ...newBook, title: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <input
          placeholder="Author"
          value={newBook.author}
          onChange={(e) => setNewBook({ ...newBook, author: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <input
          placeholder="Description"
          value={newBook.description}
          onChange={(e) => setNewBook({ ...newBook, description: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <input
          placeholder="Release Date"
          value={newBook.releaseDate}
          onChange={(e) => setNewBook({ ...newBook, releaseDate: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <button type="submit" className="w-full p-2 text-white bg-blue-500 rounded hover:bg-blue-600">
          Add Book
        </button>
      </form>

      {/* Update a book */}
      <form onSubmit={handleUpdateBook} className="mb-6 p-4 bg-blue-100 rounded shadow">
        <input
          placeholder="ID"
          value={updateBook.id}
          onChange={(e) => setUpdateBook({ ...updateBook, id: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <input
          placeholder="Title"
          value={updateBook.title}
          onChange={(e) => setUpdateBook({ ...updateBook, title: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <input
          placeholder="Author"
          value={updateBook.author}
          onChange={(e) => setUpdateBook({ ...updateBook, author: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <input
          placeholder="Description"
          value={updateBook.description}
          onChange={(e) => setUpdateBook({ ...updateBook, description: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <input
          placeholder="Release Date"
          value={updateBook.releaseDate}
          onChange={(e) => setUpdateBook({ ...updateBook, releaseDate: e.target.value })}
          className="mb-2 w-full p-2 border border-gray-300 rounded"
        />
        <button type="submit" className="w-full p-2 text-white bg-blue-500 rounded hover:bg-blue-600">
          Update Book
        </button>
      </form>

      {/* Display books */}
      <div className="space-y-4">
        {books.map((book) => (
          <div key={book.ID} className="flex items-center justify-between bg-white p-4 rounded-lg shadow">
            <CardComponent card={book} />
            <button onClick={() => deleteBook(book.ID)} className={`${btnColor} text-white py-2 px-4 rounded`}>
              Delete Book
            </button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default UserInterface;
