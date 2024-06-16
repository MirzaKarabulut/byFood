import React from 'react';

interface Card {
  id: number;
  title: string;
  author: string;
  description: string;
  releaseDate: string;
}

const CardComponent: React.FC<{ card: Card }> = ({ card }) => {
  return (
    <div className="bg-white shadow-lg rounded-lg p-2 mb-2 hover:bg-gray-100">
      <div className="text-sm text-gray-600">{card.id}</div>
      <div className="text-lg font-semibold text-gray-800">{card.title}</div>
      <div className="text-md text-gray-700">{card.author}</div>
      <div className="text-md text-gray-700">{card.description}</div>
      <div className="text-md text-gray-700">{card.releaseDate}</div>
    </div>
  );
}

export default CardComponent;