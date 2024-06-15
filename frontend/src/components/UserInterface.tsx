import React, { useState, useEffect } from "react";
import axios from "axios";
import CardComponent from "./CardComponent";

interface Book {
  id: number;
  title: string;
  description: string;
}

interface UserInterfaceProps {
  backendName: string;
}

const UserInterface: React.FC<UserInterfaceProps> = ({ backendName }) => {
  const apiURL = "http://localhost:8080";
  const [books, setBooks] = useState<Book[]>([]);
  const [newBook, setNewBook] = useState({ title: "", description: "" });
  const [updateBook, setUpdateBook] = useState({ id: "", title: "", description: "" });
  
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
        console.error("Error fetching data", error);
      }
    };
    fetchData();
  }, [backendName, apiURL]);


  // Delete a book
  const deleteBook = async (userId: number) => {
    try {
      await axios.delete(`${apiURL}/api/${backendName}/books/${booksId}`);
      setBooks(books.filter((book) => book.id !== userId));
    } catch (error) {
      console.error('Error deleting book:', error);
    }
  }
  return (
    <div className={`user-interface ${bgColor} ${backendName} w-full max-w-md p-4 my-4 rounded shadow`}>
       <h2 className="text-xl font-bold text-center text-white mb-6">{`${backendName.charAt(0).toUpperCase() + backendName.slice(1)} Backend`}</h2>
      
      {/* Display books */}
      <div className="space-y-4">
        {books.map((book) => (
          <div key={book.id} className="flex items-center justify-between bg-white p-4 rounded-lg shadow">
            <CardComponent card={book} />
            <button onClick={() => deleteBook(book.id)} className={`${btnColor} text-white py-2 px-4 rounded`}>
              Delete Book
            </button>
          </div>
        ))}
      </div>

    </div>

      );
};

export default UserInterface;