import React from "react";

interface Card {
  id: number;
  title: string;
  description: string;
}

const CardComponent: React.FC<{card: Card} > = ({ card}) => {
  return (
    <div className="bg-white shadow-lg rounded-lg p-2m mb-2 hover:bg-gray-100">
      <h1 className="text-sm text-gray-600">{card.id}</h1>
      <h2 className="text-lg font-semibold text-gray-800">{card.title}</h2>
      <p className="text-md text-gray-700">{card.description}</p>
    </div>
  );
}

export default CardComponent;